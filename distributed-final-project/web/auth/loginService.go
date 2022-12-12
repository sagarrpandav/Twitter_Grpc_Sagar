package auth

import (
	authGrpc "distributed-final-project/web/gen/auth"
	"distributed-final-project/web/globals"
	grpcClientFactory "distributed-final-project/web/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginUserCred := globals.LoginUser{}

		if err := c.BindJSON(&loginUserCred); err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		serviceDial := grpcClientFactory.GetGRPCService("loginService")
		service := authGrpc.NewAuthClient(serviceDial)

		loginUserProto := authGrpc.LoginUser{
			Email:    loginUserCred.Email,
			Password: loginUserCred.Password,
		}

		user, err := service.SignIn(c, &loginUserProto)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "Error Logging In",
			})
			return
		} else {
			c.SetCookie("userCookie", user.UserHash, 60, "/", "", false, true)
			user.UserHash = ""
			c.JSON(http.StatusOK, user)
		}
	}
}

func signUpHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		newUser := globals.User{}
		if err := c.BindJSON(&newUser); err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		serviceDial := grpcClientFactory.GetGRPCService("loginService")
		service := authGrpc.NewAuthClient(serviceDial)
		protoUser := authGrpc.User{
			Id:        0,
			FirstName: newUser.FirstName,
			LastName:  newUser.LastName,
			Email:     newUser.Email,
			Password:  newUser.Password,
		}
		user, err := service.SignUp(c, &protoUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "Error Signing Up",
			})
			return
		} else {
			c.SetCookie("userCookie", user.UserHash, 60, "/", "localhost", false, true)
			user.UserHash = ""
			c.JSON(http.StatusOK, user)
		}
	}
}

func LogoutPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("userCookie")

		serviceDial := grpcClientFactory.GetGRPCService("loginService")
		service := authGrpc.NewAuthClient(serviceDial)

		userToken := authGrpc.Token{
			UserHash: cookie,
		}

		res, _ := service.SignOut(c, &userToken)

		c.SetCookie("userCookie", res.GetMessage(), -1, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{
			"Success": "User Logged Out",
		})
	}
}
