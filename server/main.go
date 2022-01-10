package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/serdyanuk/grpc-example/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type service struct {
	pb.UnimplementedProductsServer

	products map[string]*pb.Product
}

func (s *service) AddProduct(ctx context.Context, product *pb.Product) (*pb.ProductID, error) {
	product.Id = uuid.NewString()
	s.products[product.Id] = product

	return &pb.ProductID{Value: product.Id}, nil
}

func (s *service) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	// meta data example
	if data, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Println(data.Get("time"))
	}

	// get product by product id
	if p, ok := s.products[in.Value]; ok {
		return p, nil
	}

	return nil, fmt.Errorf("product not found %s", in.Value)
}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(ProductInterceptor))
	pb.RegisterProductsServer(s, &service{
		products: make(map[string]*pb.Product),
	})
	fmt.Println("Server started on port 3000")
	log.Fatal(s.Serve(lis))
}

func ProductInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println("time: ", time.Now())
	fmt.Println("method: ", info.FullMethod)
	msg, err := handler(ctx, req)
	fmt.Println("message:", msg)
	return msg, err
}
