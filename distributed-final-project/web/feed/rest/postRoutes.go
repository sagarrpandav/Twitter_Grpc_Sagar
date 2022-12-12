package rest

import (
	"github.com/gin-gonic/gin"
)

func PostPrivateRoutes(g *gin.RouterGroup) {
	g.GET("/users", getUserHandler())
	g.POST("/addPost", addPostHandler())
	g.GET("/getPost/:userId", getPostHandler())
}
