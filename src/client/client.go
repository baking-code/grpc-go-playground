package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/baking-code/grpc-go-playground/src/recipes"
	"google.golang.org/grpc"
)

const (
	address   = "localhost:50051"
	defaultId = "1"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewRecipesClient(conn)

	id := defaultId
	if len(os.Args) > 1 {
		id = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetRecipe(ctx, &pb.GetRecipeRequest{Id: id})
	if err != nil {
		log.Fatalf("Could not get recipe: %v", err)
	}
	log.Printf("Recipe: %s", r)
}
