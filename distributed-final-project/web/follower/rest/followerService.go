package rest

import (
	followGrpc "distributed-final-project/web/gen/follower"
	grpcClientFactory "distributed-final-project/web/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FollowUserHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		serviceDial := grpcClientFactory.GetGRPCService("followService")
		service := followGrpc.NewFollowerClient(serviceDial)

		followUserProto := followGrpc.FollowRequest{
			SelfId:      int32(100),
			OtherUserId: int32(101),
			Follow:      true,
		}

		res, err := service.FollowUser(context, &followUserProto)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"Error": "Error Logging In",
			})
			return
		} else {
			context.JSON(http.StatusOK, res)
		}
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
