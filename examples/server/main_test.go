package main

import (
	"context"
	"log"
	"net"
	"testing"

	userpb "github.com/anuj070894/go_microservices_new/gen/go/users/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func newServer(t *testing.T, register func(srv *grpc.Server)) *grpc.ClientConn {
	// server initialization
	list := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		list.Close()
	})

	srv := grpc.NewServer()
	t.Cleanup(func() { srv.Stop() })

	register(srv)

	// svc := userService{}

	// userpb.RegisterUserServiceServer(srv, &svc)

	go func() {
		if err := srv.Serve(list); err != nil {
			log.Fatalf("srv.serve error: %v", err)
		}
	}()

	// test
	dialer := func(context.Context, string) (net.Conn, error) {
		return list.Dial()
	}

	conn, err := grpc.DialContext(context.Background(), "", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	t.Cleanup(func() { conn.Close() })
	// conn.Close()
	if err != nil {
		t.Fatalf("grpc.DialContext: %v", err)
	}

	return conn
}
func TestUserService_GetUser(t *testing.T) {
	svc := userService{}
	conn := newServer(t, func(srv *grpc.Server) {
		userpb.RegisterUserServiceServer(srv, &svc)
	})

	client := userpb.NewUserServiceClient(conn)
	res, err := client.GetUser(context.Background(), &userpb.GetUserRequest{
		Uuid: "123",
	})

	if err != nil {
		t.Fatalf("Error getting user %v", err)
	}

	if res.User.Uuid != "123" && res.User.FullName != "Anuj" {
		t.Fatalf("Unexpected values %v", res.User.Uuid)
	}
}
