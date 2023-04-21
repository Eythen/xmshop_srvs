package handler

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"xmshop_srvs/goods_srv/global"
	"xmshop_srvs/goods_srv/model"
	"xmshop_srvs/goods_srv/proto"
)

// 商品分类
func (g *GoodsServer) GetAllCategorysList(context.Context, *emptypb.Empty) (*proto.CategoryListResponse, error) {
	var categories []model.Category
	global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categories)
	b, _ := json.Marshal(&categories)
	return &proto.CategoryListResponse{JsonData: string(b)}, nil
}

// 获取子分类
func (g *GoodsServer) GetSubCategory(c context.Context, req *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	var category model.Category
	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}

	categoryListResponse := proto.SubCategoryListResponse{}
	categoryListResponse.Info = &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		ParentCategory: *category.ParentCategoryID,
		Level:          category.Level,
		IsTab:          category.IsTab,
	}

	var subCategories []model.Category
	var subCategoriesResponse []*proto.CategoryInfoResponse
	preloads := "SubCategory"
	if category.Level == 1 {
		preloads = "SubCategory.SubCategory"
	}

	global.DB.Where(&model.Category{ParentCategoryID: &req.Id}).Preload(preloads).Find(&subCategories)
	for _, category := range subCategories {
		subCategoriesResponse = append(subCategoriesResponse, &proto.CategoryInfoResponse{
			Id:             category.ID,
			Name:           category.Name,
			ParentCategory: *category.ParentCategoryID,
			Level:          category.Level,
			IsTab:          category.IsTab,
		})
	}
	categoryListResponse.SubCategorys = subCategoriesResponse
	return &categoryListResponse, nil
}

func (g *GoodsServer) CreateCategory(c context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	category := model.Category{}

	category.Name = req.Name
	category.Level = req.Level
	if category.Level != 1 {
		//查询父级是否存在，也可以交由前端去查询后再调用该接口
		category.ParentCategoryID = &req.ParentCategory
	}
	category.IsTab = req.IsTab

	global.DB.Save(&category)

	return &proto.CategoryInfoResponse{Id: category.ID}, nil
}

func (g *GoodsServer) DeleteCategory(c context.Context, req *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Category{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	return &emptypb.Empty{}, nil
}

func (g *GoodsServer) UpdateCategory(c context.Context, req *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	var category model.Category

	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.ParentCategory != 0 {
		category.ParentCategoryID = &req.ParentCategory
	}
	if req.Level != 0 {
		category.Level = req.Level
	}
	if req.IsTab {
		category.IsTab = req.IsTab
	}

	global.DB.Save(&category)
	return &emptypb.Empty{}, nil
}
