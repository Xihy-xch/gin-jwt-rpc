package main

import (
	"jwt"
	middle "middleWare"
	"github.com/gin-gonic/gin"
)

var secret = "secret"

var jwtInfo *jwt.Jwt

func main() {
	jwtInfo = jwt.NewJwt()
	jwtInfo.SetSecret(secret)
	route := gin.Default()
	route.POST("/login", setJwtInfo, middle.HandlerJwt)

	route.Run(":8080")
}
func setJwtInfo(c *gin.Context) {
	c.Set("jwtInfo", jwtInfo)
	c.Next()
}