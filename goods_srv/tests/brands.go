package main

import (
	"context"
	"fmt"
	"xmshop_srvs/goods_srv/proto"
)

func TestGetBrandList() {
	rsp, err := client.BrandList(context.Background(), &proto.BrandFilterRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, brand := range rsp.Data {
		fmt.Println(brand.Name)
	}
}
