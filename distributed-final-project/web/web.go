package main

import (
	loginRoutes "distributed-final-project/web/auth"
	auth "distributed-final-project/web/auth/middleware"
	loginService "distributed-final-project/web/auth/server"
	"distributed-final-project/web/globals"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	go loginService.CreateServer()
	router := gin.Default()
	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	loginGroup := router.Group("/")
	loginGroup.Use(auth.AuthRequired)
	loginRoutes.LoginPublicRoutes(loginGroup)

	router.Run("localhost:8080")
}

