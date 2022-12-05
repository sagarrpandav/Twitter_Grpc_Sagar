package login

import (
	authImpl "distributed-final-project/web/auth/gRPCImpl"
	pb "distributed-final-project/web/gen/auth"
	"google.golang.org/grpc"
	"log"
	"net"
)

func CreateServer() {
	listener, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServer(grpcServer, &authImpl.AuthApiServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Auth GRPC Error")
	}
}