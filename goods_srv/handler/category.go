package handler

import (
	"context"
	"encoding/json"
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

//// 获取子分类
//GetSubCategory(context.Context, *CategoryListRequest) (*SubCategoryListResponse, error)
//CreateCategory(context.Context, *CategoryInfoRequest) (*CategoryInfoResponse, error)
//DeleteCategory(context.Context, *DeleteCategoryRequest) (*emptypb.Empty, error)
//UpdateCategory(context.Context, *CategoryInfoRequest) (*emptypb.Empty, error)
