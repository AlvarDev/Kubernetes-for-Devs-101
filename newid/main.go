package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	pb "newidservice/pb"

	"google.golang.org/grpc"
)

const (
	port = ":50052"
	max  = 1000
	min  = 2
)

// server is used to implement RPC server
type server struct {
	pb.UnimplementedReminderServiceServer
}

func (s *server) GetNewId(ctx context.Context, req *pb.GetNewIdRequest) (*pb.GetNewIdResponse, error) {
	fmt.Println("Request received!")
	newId := rand.Intn(max-min) + min

	return &pb.GetNewIdResponse{
		Remainder: &pb.Remainder{
			RemainderId: int64(newId),
			Remainder:   req.GetRemainder()},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listening: %v:", err)
	}

	s := grpc.NewServer()
	pb.RegisterReminderServiceServer(s, &server{})
	fmt.Printf("server listening at: %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
