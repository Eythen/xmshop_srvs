package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"xmshop_srvs/user_srv/proto"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	userClient = proto.NewUserClient(conn)
}

func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    0,
		PSize: 2,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.NickName, user.Password)
		checkRsp, err := userClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
			Password:          "admin123",
			EncryptedPassword: user.Password,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkRsp.Success)
	}
}

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			NickName: fmt.Sprintf("bobby%d", i),
			Password: "admin123",
			Mobile:   fmt.Sprintf("1882000724%d", i),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}

func TestGetUserByMobile() {
	rsp, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: "18782222220"})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func TestGetUserById() {
	rsp, err := userClient.GetUserById(context.Background(), &proto.IdRequest{Id: 1})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func TestUpdateUser() {
	rsp, err := userClient.UpdateUser(context.Background(), &proto.UpdateUserInfo{
		Id:       2,
		NickName: "cesf",
		Gender:   "female",
		Birthday: 1681108332,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func main() {
	Init()
	TestGetUserList()
	//TestCreateUser()
	//TestGetUserByMobile()
	//TestGetUserById()
	//TestUpdateUser()
	defer conn.Close()
}
