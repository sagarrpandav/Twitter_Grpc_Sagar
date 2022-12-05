package middleware

import (
	authGrpc "distributed-final-project/web/gen/auth"
	grpcClientFactory "distributed-final-project/web/util"
	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {
	cookie, _ := c.Cookie("userCookie")
	serviceDial := grpcClientFactory.GetGRPCService("loginService")
	service := authGrpc.NewAuthClient(serviceDial)
	userToken := authGrpc.Token{
		UserHash: cookie,
	}
	_, err := service.ValidateUserLoggedIn(c, &userToken)
	if err != nil {
		c.Next()
		return
	}
	c.Next()
}
