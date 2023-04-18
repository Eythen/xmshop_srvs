package main

import (
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

func main() {
	Init()
	TestGetSubCategoryList()
	defer conn.Close()
}
