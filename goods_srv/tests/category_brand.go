package main

import (
	"context"
	"fmt"
	"xmshop_srvs/goods_srv/proto"
)

func TestCategoryBrandList() {
	rsp, err := client.CategoryBrandList(context.Background(), &proto.CategoryBrandFilterRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.Data)
}

func TestGetCategoryBrandList() {
	rsp, err := client.GetCategoryBrandList(context.Background(), &proto.CategoryInfoRequest{
		Id: 135475,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.Data)
}
