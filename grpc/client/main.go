package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ivandhitya/sinau/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// membuat koneksi baru ke server gRPC
	conn, err := grpc.NewClient(":3333", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// membuat client untuk mengakses user service
	c := service.NewUserServiceClient(conn)

	// memeprsiapkan data request pada fucntion GetUserInfo
	req := &service.UserRequest{UserId: 123}

	// request ke server GRPC
	res, err := c.GetUserInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get user info: %v", err)
	}
	// menampilkan response dari endpoint/ function GetUserInfo
	fmt.Printf("User Info: %v\n", res)
}
