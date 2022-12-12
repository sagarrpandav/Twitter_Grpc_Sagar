package server

import (
	feedImpl "distributed-final-project/web/feed/gRPCImpl"
	pb "distributed-final-project/web/gen/feed"
	"google.golang.org/grpc"
	"log"
	"net"
)

func CreateFeedServer() {
	listener, err := net.Listen("tcp", "localhost:8091")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterFeedServer(grpcServer, &feedImpl.FeedApiServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Auth GRPC Error")
	}
}
