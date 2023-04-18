package main

import (
	"context"
	"fmt"
	"xmshop_srvs/goods_srv/proto"
)

func TestGoodsList() {
	rsp, err := client.GoodsList(context.Background(), &proto.GoodsFilterRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.Data)
}
