package main

import (
	"fmt"
	"jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var secret = "secret"

func main() {
	route := gin.Default()

	route.POST("/login", func(c *gin.Context) {
		jwtCookie, err := c.Cookie("jwtCookie")

		if err != nil {
			jwtCookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", jwtCookie)
		// c.String(http.StatusOK, "Hello vister!")
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "admin" && password == "admin" {
			jwtInfo := jwt.NewJwt()
			jwtInfo.SetHeader("HS256")
			jwtInfo.SetPayload(username)
			header := jwtInfo.GetHeaderToBase64()
			payload := jwtInfo.GetPayloadToBase64()
			jwtInfo.SetSecret(secret)
			signature := jwtInfo.ComputeHmacSha256(header + "." + payload)
			token := header + "." + payload + "." + signature

			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})

		}

	})

	route.Run(":8080")
}
