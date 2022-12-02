package login

import (
	"distributed-final-project/web/globals"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	//"log"
	"net/http"
)

func LoginPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Already Signed In",
			})
			return
		}
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "Already Signed In",
			})
			return
		}
	}
}

func signUpHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//session := sessions.Default(c)
		//user := session.Get(globals.Userkey)
		newUser := globals.User{}
		if err := c.BindJSON(&newUser); err!=nil{
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(newUser)
		var users = globals.Users
		for _, x := range users{
			if x.Email == newUser.Email {
				c.JSON(http.StatusOK, gin.H{
					"Error": "User ID already exists",
				})
				return
			}
		}
		globals.Users = append(globals.Users, newUser)
	}
}