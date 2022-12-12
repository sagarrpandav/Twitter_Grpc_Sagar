package rest

import (
	feedGrpc "distributed-final-project/web/gen/feed"
	"distributed-final-project/web/globals"
	grpcClientFactory "distributed-final-project/web/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

		serviceDial := grpcClientFactory.GetGRPCService("feedService")
		service := feedGrpc.NewFeedClient(serviceDial)

		protoUser := feedGrpc.UserId{
			Id: int32(userId),
		}

		_, err = service.GetPost(c, &protoUser)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"Success": "Post Added Successfully",
			})
		}
	}
}

func addPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		newPost := globals.Post{}

		if err := c.BindJSON(&newPost); err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		serviceDial := grpcClientFactory.GetGRPCService("feedService")
		service := feedGrpc.NewFeedClient(serviceDial)

		protoPost := feedGrpc.Post{
			PostId:     newPost.PostId,
			UserId:     newPost.UserId,
			Content:    newPost.Content,
			DatePosted: newPost.DatePosted,
		}

		_, err := service.AddPost(c, &protoPost)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"Success": "Post Added Successfully",
			})
		}
	}
}

func getUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		serviceDial := grpcClientFactory.GetGRPCService("feedService")
		service := feedGrpc.NewFeedClient(serviceDial)

		users, _ := service.GetUsers(c, nil)

		c.JSON(http.StatusOK, users)
	}
}
