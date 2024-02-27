package main

import (
	"context"
	"fmt"
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

	stream, err := client.ConsumeBeatsPerMinute(context.Background())

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

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalln("Closing err ", err)
	}

	fmt.Println("Total messages", res.GetTotal())
}
