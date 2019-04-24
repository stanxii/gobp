
//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "downvideo/grpc/proto"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {}

// implement gprc inteface
func (s *server) DownloadVideo(ctx context.Context, in *pb.DVRequest) (*pb.DVReply, error) {
    log.Printf( "Recived from grpc clent msg::: %v", in.Name)
	return &pb.DVReply{Message: "Hello " + in.Name}, nil
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDownVideoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
