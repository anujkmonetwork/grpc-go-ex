package main

import (
	"context"
	"fmt"
	"log"
	"time"

	userpb "github.com/anuj070894/go_microservices_new/gen/go/users/v1"
	wearablepb "github.com/anuj070894/go_microservices_new/gen/go/wearable/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(UnaryClientInterceptor),
		grpc.WithStreamInterceptor(StreamClientInterceptor),
	}

	conn, err := grpc.Dial("localhost:9879", opts...)

	if err != nil {
		log.Fatalf("Fail to connect: %v", err)
	}

	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()
	res, err := client.GetUser(ctx, &userpb.GetUserRequest{
		Uuid: "123",
	})

	if err != nil {
		log.Fatalf("Fail to get user: %v", err)
	}

	fmt.Printf("%+v\n", res)

	client2 := wearablepb.NewWearableServiceClient(conn)

	stream, err := client2.ConsumeBeatsPerMinute(context.Background())

	if err != nil {
		log.Fatalln("Consuming stream error: ", err)
	}

	for i := 0; i < 10; i++ {
		err := stream.Send(&wearablepb.ConsumeBeatsPerMinuteRequest{
			Uuid:   "Anuj",
			Value:  100,
			Minute: uint32(i),
		})

		if err != nil {
			log.Fatalln("Sending value error ", err)
		}

		time.Sleep(100 * time.Millisecond)
	}

	res2, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalln("Closing err ", err)
	}

	fmt.Println("Total messages", res2.GetTotal())
}
