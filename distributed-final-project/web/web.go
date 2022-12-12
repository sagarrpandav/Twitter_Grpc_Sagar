package main

import (
	loginRoutes "distributed-final-project/web/auth"
	auth "distributed-final-project/web/auth/middleware"
	loginGRPC "distributed-final-project/web/auth/server"
	feedRoutes "distributed-final-project/web/feed/rest"
	feedGRPC "distributed-final-project/web/feed/server"
	followerGRPC "distributed-final-project/web/follower/server"
	"distributed-final-project/web/globals"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	go loginGRPC.CreateServer()
	go feedGRPC.CreateFeedServer()
	go followerGRPC.CreateFollowerServer()

	router := gin.Default()
	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = 12 * time.Hour
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(corsConfig))

	loginGroup := router.Group("/")
	loginGroup.Use(auth.AuthRequired)
	loginRoutes.LoginPublicRoutes(loginGroup)

	feedGroup := router.Group("/")
	feedGroup.Use(auth.AuthRequired)
	feedRoutes.PostPrivateRoutes(feedGroup)

	postGroup := router.Group("/")
	postGroup.Use(auth.AuthRequired)

	router.Run("localhost:8080")
}
