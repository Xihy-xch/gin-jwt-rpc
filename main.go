package main

import (
	"fmt"
	"jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var secret = "secret"


var jwtInfo *jwt.Jwt

func main() {
	jwtInfo = jwt.NewJwt()
	jwtInfo.SetSecret(secret)
	route := gin.Default()
	route.POST("/login", func(c *gin.Context) {
		jwtCookie, _ := c.Cookie("jwtCookie")
		username := c.PostForm("username")
		if len(jwtCookie) > 0 && jwtInfo.CheckToken(jwtCookie) {
			token := jwtInfo.GetSecrectToken(username)
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
			fmt.Println("token success!",token)
		} else {
			password := c.PostForm("password")
			if username == "admin" && password == "admin" {
				
				token := jwtInfo.GetSecrectToken(username)
				fmt.Println("username and password is correct!", token)
				c.JSON(http.StatusOK, gin.H{
					"token": token,
				})

			} else {
				c.String(http.StatusForbidden, "username or password is wrong!")
			}
		} 
		// c.String(http.StatusOK, "Hello vister!")

	})

	route.Run(":8080")
}
