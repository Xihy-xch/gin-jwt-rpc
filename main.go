package main

import (
	"github.com/gin-gonic/gin"
	"workspace/gin-jwt-rpc/jwt"
	"workspace/gin-jwt-rpc/middleWare"
	"workspace/gin-jwt-rpc/sql"
)

var secret = "secret"

var jwtInfo *jwt.Jwt

func main() {
	db := sql.Connection("root", "root", "gin_jwt_rpc")
	defer db.Close()
	jwtInfo = jwt.NewJwt()
	jwtInfo.SetSecret(secret)
	route := gin.Default()
	route.POST("/login", setJwtInfo, middleWare.HandlerJwt)

	route.Run(":8080")
}
func setJwtInfo(c *gin.Context) {
	c.Set("jwtInfo", jwtInfo)
	c.Next()
}
