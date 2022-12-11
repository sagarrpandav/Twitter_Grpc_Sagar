package rest

import (
	grpcClientFactory "distributed-final-project/web/util"
	"github.com/gin-gonic/gin"
)

func FollowUserHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		grpcClientFactory.GetGRPCService('followService')
	}
}

func UnfollowUser() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func GetUsersHandler() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
