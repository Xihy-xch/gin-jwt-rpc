package main

import (
	"jwt"
	middle "middleWare"
	"github.com/gin-gonic/gin"
	"sql"
)

var secret = "secret"

var jwtInfo *jwt.Jwt

func main() {
	db := sql.Connection("root", "root", "gin_jwt_rpc")
	defer db.Close()
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