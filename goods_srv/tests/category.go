package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"xmshop_srvs/goods_srv/proto"
)

func TestGetCategoryList() {
	rsp, err := client.GetAllCategorysList(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.JsonData)
}

func TestGetSubCategoryList() {
	rsp, err := client.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: 130358,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.SubCategorys)
}
