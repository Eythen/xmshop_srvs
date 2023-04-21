package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"xmshop_srvs/goods_srv/global"
	"xmshop_srvs/goods_srv/model"
	"xmshop_srvs/goods_srv/proto"
)

type GoodsServer struct {
	*proto.UnimplementedGoodsServer
}

func ModelToResponse(goods model.Goods) proto.GoodsInfoResponse {
	return proto.GoodsInfoResponse{
		Id:              goods.ID,
		CategoryId:      goods.CategoryID,
		Name:            goods.Name,
		GoodsSn:         goods.GoodsSn,
		ClickNum:        goods.ClickNum,
		SoldNum:         goods.SoldNum,
		FavNum:          goods.FavNum,
		MarketPrice:     goods.MarketPrice,
		ShopPrice:       goods.ShopPrice,
		GoodsBrief:      goods.GoodsBrief,
		ShipFree:        goods.ShipFree,
		Images:          goods.Images,
		DescImages:      goods.DescImages,
		GoodsFrontImage: goods.GoodsFrontImage,
		IsNew:           goods.IsNew,
		IsHot:           goods.IsHot,
		OnSale:          goods.OnSale,
		Category: &proto.CategoryBriefInfoResponse{
			Id:   goods.Category.ID,
			Name: goods.Category.Name,
		},
		Brand: &proto.BrandInfoResponse{
			Id:   goods.Brands.ID,
			Name: goods.Brands.Name,
			Logo: goods.Brands.Logo,
		},
	}
}

// 商品接口
func (g *GoodsServer) GoodsList(c context.Context, req *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
	//关键词搜索、查询新品、查询热门商品、通过价格区间筛选、通过商品分类筛选
	goodsListResponse := &proto.GoodsListResponse{}

	var goods []model.Goods
	localDB := global.DB.Model(&model.Goods{})
	if req.KeyWords != "" {
		localDB = localDB.Where("name LIKE ?", "%"+req.KeyWords+"%")
	}
	if req.IsHot {
		localDB = localDB.Where(model.Goods{IsHot: true})
	}
	if req.IsNew {
		localDB = localDB.Where("is_new=true")
	}

	if req.PriceMin > 0 {
		localDB = localDB.Where("shop_price >= ?", req.PriceMin)
	}
	if req.PriceMax > 0 {
		localDB = localDB.Where("shop_price <= ?", req.PriceMax)
	}

	if req.Brand > 0 {
		localDB = localDB.Where("brand_id = ?", req.Brand)
	}

	//通过category去查询商品
	var subQuery string
	if req.TopCategory > 0 {
		var category model.Category
		if result := global.DB.First(&category, req.TopCategory); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "商品分类不存在")
		}

		if category.Level == 1 {
			subQuery = fmt.Sprintf("select id from category where parent_category_id in (select id from category where parent_category_id=%d)", req.TopCategory)
		} else if category.Level == 2 {
			subQuery = fmt.Sprintf("select id from category where parent_category_id where parent_category_id=%d", req.TopCategory)
		} else if category.Level == 3 {
			subQuery = fmt.Sprintf("select id from category where parent_category_id where id=%d", req.TopCategory)
		}
		localDB = localDB.Where(fmt.Sprintf("category_id in (%s)", subQuery))
	}

	var count int64
	localDB.Count(&count)
	goodsListResponse.Total = int32(count)

	result := localDB.Preload("Category").Preload("Brands").Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&goods)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, good := range goods {
		GoodsInfoResponse := ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &GoodsInfoResponse)
	}
	return goodsListResponse, nil
}

// 现在用户提交订单有多个商品，你得批量查询商品的信息吧
func (g *GoodsServer) BatchGetGoods(c context.Context, req *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
	goodsListResponse := &proto.GoodsListResponse{}
	var goods []model.Goods

	result := global.DB.Find(&goods, req.Id)
	for _, good := range goods {
		goodsInfoResponse := ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &goodsInfoResponse)
	}
	goodsListResponse.Total = int32(result.RowsAffected)

	return goodsListResponse, nil
}

func (g *GoodsServer) GetGoodsDetail(c context.Context, req *proto.GoodInfoRequest) (*proto.GoodsInfoResponse, error) {
	var goods model.Goods
	if result := global.DB.First(&goods, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	goodsInfoResponse := ModelToResponse(goods)

	return &goodsInfoResponse, nil
}

func (g *GoodsServer) CreateGoods(c context.Context, req *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	goods := model.Goods{
		Category:        category,
		CategoryID:      req.CategoryId,
		Brands:          brand,
		BrandsID:        req.BrandId,
		OnSale:          req.OnSale,
		ShipFree:        req.ShipFree,
		IsNew:           req.IsNew,
		IsHot:           req.IsHot,
		Name:            req.Name,
		GoodsSn:         req.GoodsSn,
		MarketPrice:     req.MarketPrice,
		ShopPrice:       req.ShopPrice,
		GoodsBrief:      req.GoodsBrief,
		Images:          req.Images,
		DescImages:      req.DescImages,
		GoodsFrontImage: req.GoodsFrontImage,
	}

	global.DB.Save(&goods)
	return &proto.GoodsInfoResponse{Id: goods.ID}, nil
}

func (g *GoodsServer) DeleteGoods(c context.Context, req *proto.DeleteGoodsInfo) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Goods{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}

	return &emptypb.Empty{}, nil
}

func (g *GoodsServer) UpdateGoods(c context.Context, req *proto.CreateGoodsInfo) (*emptypb.Empty, error) {
	var goods model.Goods

	if result := global.DB.First(&goods, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}

	if req.CategoryId != 0 {
		var category model.Category
		if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
		}
		goods.Category = category
		goods.CategoryID = req.CategoryId
	}

	if req.BrandId != 0 {
		var brand model.Brands
		if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
		}
		goods.Brands = brand
		goods.BrandsID = req.BrandId
	}

	if req.Name != "" {
		goods.Name = req.Name
	}
	if req.GoodsSn != "" {
		goods.GoodsSn = req.GoodsSn
	}
	if req.MarketPrice != 0 {
		goods.MarketPrice = req.MarketPrice
	}
	if req.ShopPrice != 0 {
		goods.ShopPrice = req.ShopPrice
	}
	if req.GoodsBrief != "" {
		goods.GoodsBrief = req.GoodsBrief
	}
	if goods.ShipFree || req.ShipFree {
		goods.ShipFree = req.ShipFree
	}

	if len(req.Images) > 0 {
		goods.Images = req.Images
	}
	if len(req.DescImages) > 0 {
		goods.DescImages = req.DescImages
	}
	if req.GoodsFrontImage != "" {
		goods.GoodsFrontImage = req.GoodsFrontImage
	}

	goods.IsNew = req.IsNew
	goods.IsHot = req.IsHot
	goods.OnSale = req.OnSale

	global.DB.Save(&goods)
	return &emptypb.Empty{}, nil
}
