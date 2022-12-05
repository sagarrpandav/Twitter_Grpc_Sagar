package auth

import (
	"github.com/gin-gonic/gin"
)

func LoginPublicRoutes(g *gin.RouterGroup) {
	g.POST("/login", LoginPostHandler())
	g.GET("/logout", LogoutPostHandler())
	g.PUT("/signup", signUpHandler())
}

func PrivateRoutes(g *gin.RouterGroup) {
	//
	//g.GET("/dashboard", controllers.DashboardGetHandler())
	//g.GET("/logout", controllers.LogoutGetHandler())

}