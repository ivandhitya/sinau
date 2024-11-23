package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ivandhitya/sinau/grpc/service"
	"google.golang.org/grpc"
)

type server struct {
	service.UnimplementedUserServiceServer
}

func (s *server) GetUserInfo(ctx context.Context, req *service.UserRequest) (*service.UserResponse, error) {
	// Proses untuk mendapatkan data user
	user := &service.UserResponse{
		UserId: req.UserId,
		Name:   "Khamzat Chimaev",
		Email:  "khamzat.chimaev@ufc.com",
	}
	return user, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service.RegisterUserServiceServer(s, &server{})
	fmt.Println("Server is running at port 3333...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
