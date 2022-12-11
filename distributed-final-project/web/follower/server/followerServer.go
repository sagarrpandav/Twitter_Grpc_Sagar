package server

import (
	pb "distributed-final-project/web/gen/follower"
	"google.golang.org/grpc"
	"log"
	"net"
)

func CreateServer() {
	listener, err := net.Listen("tcp", "localhost:8092")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterFollowerServer(grpcServer, &followerI)
}
