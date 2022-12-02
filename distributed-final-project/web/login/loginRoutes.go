package login

import (
	"github.com/gin-gonic/gin"
)

func LoginPublicRoutes(g *gin.RouterGroup) {
	g.GET("/login", LoginPostHandler())
	g.GET("/logout", LoginPostHandler())
	g.PUT("/signup", signUpHandler())
}

func PrivateRoutes(g *gin.RouterGroup) {
	//
	//g.GET("/dashboard", controllers.DashboardGetHandler())
	//g.GET("/logout", controllers.LogoutGetHandler())

}