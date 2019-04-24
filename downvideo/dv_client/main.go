
package main

import (
	"context"
	"log"
	"google.golang.org/grpc"
	"os"
	"time"
	pb "downvideo/grpc/proto"
)

const (
	address  = "localhost:50051"
	defaultName = "grpc world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDownVideoClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.DownloadVideo(ctx, &pb.DVRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("Greeting: %s", r.Message)
}