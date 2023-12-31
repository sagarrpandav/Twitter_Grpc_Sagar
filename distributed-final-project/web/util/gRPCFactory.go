package util

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serviceMap = map[string]string{
	"loginService":    "localhost:8090",
	"feedService":     "localhost:8091",
	"followerService": "localhost:8092",
}

func GetGRPCService(serviceName string) *grpc.ClientConn {
	dial, err := grpc.Dial(serviceMap[serviceName], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil
	}
	return dial
}
