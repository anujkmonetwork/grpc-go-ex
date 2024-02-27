package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	wearablepb "github.com/anuj070894/go_microservices_new/gen/go/wearable/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:9879", opts...)

	if err != nil {
		log.Fatalf("failed to connect %v", err)
	}

	defer conn.Close()

	client := wearablepb.NewWearableServiceClient(conn)

	stream, err := client.CalculateBeatsPerMinute(context.Background())

	if err != nil {
		log.Fatalln("Calculating stream error: ", err)
	}

	for i := 0; i < 10; i++ {
		err := stream.Send(&wearablepb.CalculateBeatsPerMinuteRequest{
			Uuid:   "Anuj",
			Value:  uint32(2 * i),
			Minute: uint32(i),
		})

		if err != nil {
			log.Fatalln("Sending value error ", err)
		}

		time.Sleep(100 * time.Millisecond)
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalln("Close send", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("Closing", err)
		}

		fmt.Println("Total average: ", res.GetAverage())
	}
}
