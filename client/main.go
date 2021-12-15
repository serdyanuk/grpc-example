package main

import (
	"context"
	"fmt"
	"grpc3/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial(
		":3000",
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`), // This sets the initial balancing policy.
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewProductsClient(conn)
	id, err := c.AddProduct(context.TODO(), &pb.Product{
		Name:   "Foo name",
		Amount: "10.00",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("product id %s\n", id.Value)

	// meta data example
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.MD{
		"time": []string{
			time.Now().String(),
		},
	})
	product, err := c.GetProduct(ctx, &pb.ProductID{Value: id.Value})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("product %v\n", product)
}
