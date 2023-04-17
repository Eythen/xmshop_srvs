package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"xmshop_srvs/goods_srv/proto"
)

var client proto.GoodsClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client = proto.NewGoodsClient(conn)
}

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

func main() {
	Init()
	TestGetBrandList()
	defer conn.Close()
}
