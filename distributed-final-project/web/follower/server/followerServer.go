package server

import (
	followImpl "distributed-final-project/web/follower/gRPCImpl"
	pb "distributed-final-project/web/gen/follower"
	"google.golang.org/grpc"
	"log"
	"net"
)

func CreateFollowerServer() {
	listener, err := net.Listen("tcp", "localhost:8092")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterFollowerServer(grpcServer, &followImpl.FollowerServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Follower GRPC Error")
	}
}
