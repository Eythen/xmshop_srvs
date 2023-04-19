package main

import (
	"context"
	"fmt"
	"xmshop_srvs/goods_srv/proto"
)

func TestGoodsList() {
	rsp, err := client.GoodsList(context.Background(), &proto.GoodsFilterRequest{
		TopCategory: 136982,
		PriceMax:    90,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.Data)
}

func TestBatchGetGoods() {
	rsp, err := client.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: []int32{421, 422},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.Data)
}

func TestGetGoodsDetail() {
	rsp, err := client.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: 421,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}
