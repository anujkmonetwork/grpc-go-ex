package main

import (
	"context"
	"log"
	"net"
	"time"

	userpb "github.com/anuj070894/go_microservices_new/gen/go/users/v1"
	wearablepb "github.com/anuj070894/go_microservices_new/gen/go/wearable/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type userService struct{}

func (u *userService) GetUser(_ context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	time.Sleep(1 * time.Second)
	return &userpb.GetUserResponse{
		User: &userpb.User{
			Uuid:     req.Uuid,
			FullName: "Anuj",
		},
	}, nil
}

// func (u *userService) mustEmbedUnimplementedUserServiceServer() {}

func main() {
	lis, err := net.Listen("tcp", "localhost:9879")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	wearableServer := &wearableService{}
	userServer := &userService{}
	healthServer := health.NewServer()

	go func() {
		for {
			status := healthpb.HealthCheckResponse_SERVING

			// if time.Now().Second()%2 == 0 {
			// 	status = healthpb.HealthCheckResponse_NOT_SERVING
			// }

			healthServer.SetServingStatus(userpb.UserService_ServiceDesc.ServiceName, status)
			healthServer.SetServingStatus("", status)

			time.Sleep(1 * time.Second)
		}
	}()
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus(userpb.UserService_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryServerInterceptor),
		grpc.StreamInterceptor(StreamServerInterceptor),
	)
	userpb.RegisterUserServiceServer(grpcServer, userServer)
	wearablepb.RegisterWearableServiceServer(grpcServer, wearableServer)
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
