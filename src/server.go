package main

import (
	"context"
	"log"
	"net"

	"github.com/baking-code/grpc-go-playground/src/recipes"
	pb "github.com/baking-code/grpc-go-playground/src/recipes"

	"google.golang.org/grpc"
)

type instance struct {
	recipes.UnimplementedRecipesServer
}

// mustEmbedUnimplementedRecipesServer implements recipes.RecipesServer.
func (*instance) mustEmbedUnimplementedRecipesServer() {
	panic("unimplemented")
}

func (s *instance) GetRecipe(ctx context.Context, in *pb.GetRecipeRequest) (*pb.Recipe, error) {
	return &pb.Recipe{Id: "1", Name: "Beans on toast", TimeInMinutes: 8}, nil
}

const PORT = ":50051"

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	serverInstance := instance{}

	pb.RegisterRecipesServer(s, &serverInstance)

	log.Println("Server started on:", PORT)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
