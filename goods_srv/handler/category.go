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
		ParentCategory: category.ParentCategoryID,
		Level:          category.Level,
		IsTab:          category.IsTab,
	}

	var subCategories []model.Category
	var subCategoriesResponse []*proto.CategoryInfoResponse
	preloads := "SubCategory"
	if category.Level == 1 {
		preloads = "SubCategory.SubCategory"
	}

	global.DB.Where(&model.Category{ParentCategoryID: req.Id}).Preload(preloads).Find(&subCategories)
	for _, category := range subCategories {
		subCategoriesResponse = append(subCategoriesResponse, &proto.CategoryInfoResponse{
			Id:             category.ID,
			Name:           category.Name,
			ParentCategory: category.ParentCategoryID,
			Level:          category.Level,
			IsTab:          category.IsTab,
		})
	}
	categoryListResponse.SubCategorys = subCategoriesResponse
	return &categoryListResponse, nil
}

//CreateCategory(context.Context, *CategoryInfoRequest) (*CategoryInfoResponse, error)
//DeleteCategory(context.Context, *DeleteCategoryRequest) (*emptypb.Empty, error)
//UpdateCategory(context.Context, *CategoryInfoRequest) (*emptypb.Empty, error)
