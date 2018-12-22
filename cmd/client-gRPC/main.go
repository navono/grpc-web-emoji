package main

import (
	"log"
	"time"

	v1 "github.com/navono/gRPC-Web-emoji/pkg/api/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// connect to the server
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to the service: %v", err)
	}
	defer conn.Close()

	// send a request to the server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := v1.NewEmojiServiceClient(conn)
	resp, err := c.Emojize(ctx, &v1.EmojizeRequest{
		Text: "I like :pizza: and :sushi:!",
	})
	if err != nil {
		log.Fatalf("could not call service: %v", err)
	}
	log.Printf("server says: %s", resp.GetEmojizedText())
}
