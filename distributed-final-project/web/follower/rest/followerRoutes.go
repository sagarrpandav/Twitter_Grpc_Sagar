package rest

import "github.com/gin-gonic/gin"

func FollowerPrivateRoutes(g *gin.RouterGroup) {
	g.POST("/follow", FollowUserHandler())
}
