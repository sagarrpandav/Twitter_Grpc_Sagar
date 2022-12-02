package main

import (
	"distributed-final-project/web/globals"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	loginRoutes "distributed-final-project/web/login"
)

func main() {
	router := gin.Default()
	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	loginGroup := router.Group("/")
	loginRoutes.LoginPublicRoutes(loginGroup)

	router.Run("localhost:8080")
}

