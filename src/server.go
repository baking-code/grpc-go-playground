package main

import (
	"context"
	"log"
	"net"

	"github.com/baking-code/grpc-go-playground/src/greet"
	pb "github.com/baking-code/grpc-go-playground/src/greet"

	"google.golang.org/grpc"
)

type instance struct {
	greet.UnimplementedGreeterServer
}

// mustEmbedUnimplementedGreeterServer implements greet.GreeterServer.
func (*instance) mustEmbedUnimplementedGreeterServer() {
	panic("unimplemented")
}

func (s *instance) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello, " + in.Name}, nil
}

const PORT = ":50051"

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	greeterInstance := instance{}

	pb.RegisterGreeterServer(s, &greeterInstance)

	log.Println("Server started on:", PORT)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
