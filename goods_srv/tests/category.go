package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestGetCategoryList() {
	rsp, err := client.GetAllCategorysList(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.JsonData)
}
