package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/ivandhitya/sinau/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	var middlewares []grpc.ServerOption
	middlewares = append(middlewares, grpc.UnaryInterceptor(AuthInterceptor))

	s := grpc.NewServer(middlewares...)
	service.RegisterUserServiceServer(s, &server{})
	fmt.Println("Server is running at port 3333...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func AuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	// Extract token dari metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Missing metadata")
	}

	// Memeriksa token
	authHeader := md["authorization"]
	if len(authHeader) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "Authorization token is missing")
	}

	token := authHeader[0]
	// Verifikasi token
	if token != "Bearer token_rahasia" {
		return nil, errors.New("Invalid Token")
	}

	// Jika token valid, lanjutkan dengan eksekusi handler
	return handler(ctx, req)
}
