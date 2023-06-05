package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/mohfahrur/interop-service-a/grpc/item"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "Ravi"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Item to get")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewItemServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetItem(ctx, &pb.ItemRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not get item: %v", err)
	}
	log.Printf("Item: %s", r.GetMessage())
}
