package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/macario12/GRCPF2/protos"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGameServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) AddGame(ctx context.Context, in *pb.GameRequest) (*pb.GameResponse, error) {
	log.Printf("Received: %v", in.GetGameId())
	return &pb.GameResponse{Message: "Hello "}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGameServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
