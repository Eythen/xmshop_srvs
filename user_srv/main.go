package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"xmshop_srvs/user_srv/handler"
	"xmshop_srvs/user_srv/proto"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "IP address")
	Port := flag.Int("port", 50051, "Port number")

	flag.Parse()
	fmt.Println("ip:", *IP)
	fmt.Println("port:", *Port)

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = server.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
