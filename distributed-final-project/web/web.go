package main

import (
	"bytes"
	loginRoutes "distributed-final-project/web/auth"
	auth "distributed-final-project/web/auth/middleware"
	loginGRPC "distributed-final-project/web/auth/server"
	feedRoutes "distributed-final-project/web/feed/rest"
	feedGRPC "distributed-final-project/web/feed/server"
	followerRoutes "distributed-final-project/web/follower/rest"
	followerGRPC "distributed-final-project/web/follower/server"
	"distributed-final-project/web/globals"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	go loginGRPC.CreateServer()
	go feedGRPC.CreateFeedServer()
	go followerGRPC.CreateFollowerServer()

	tmpOPbj := globals.Auth{AuthDb: "[{'userId':'abc', 'token': 'jsfhsdgfhds'}, {'userId':'abc', 'token': 'jsfhsdgfhds'}]"}

	tmpOPbj1, _ := json.Marshal(tmpOPbj)
	res, err := http.Post("http://localhost:11000/key", "application/json", bytes.NewBuffer(tmpOPbj1))
	fmt.Println(err)
	fmt.Println(res)
	go globals.PostCallRaftServers(tmpOPbj)
	go globals.GetCallRaftServers("AuthDb")

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

	followerGroup := router.Group("/")
	followerGroup.Use(auth.AuthRequired)
	followerRoutes.FollowerPrivateRoutes(followerGroup)

	router.Run("localhost:8080")
}
