package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"xmshop_srvs/user_srv/handler"
	"xmshop_srvs/user_srv/initialize"
	"xmshop_srvs/user_srv/proto"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "IP address")
	Port := flag.Int("port", 50051, "Port number")

	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	flag.Parse()
	zap.S().Info("ip:", *IP)
	zap.S().Info("port:", *Port)

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	err = server.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
