package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/mohfahrur/interop-service-a/grpc/item"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

var userItem = map[string]string{
	"Rizky":     "PC",
	"Fido":      "Mouse",
	"Bhatinden": "Keyboard",
	"David":     "Monitor",
	"Ravi":      "Power Supply"}

type server struct {
	pb.UnimplementedItemServiceServer
}

func (s *server) GetItem(ctx context.Context, in *pb.ItemRequest) (*pb.ItemReply, error) {
	log.Printf("Received: %v", in.GetName())
	log.Println()
	return &pb.ItemReply{Message: userItem[in.GetName()]}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterItemServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
