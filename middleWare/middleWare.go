package middleWare

import (
	"fmt"
	"net/http"

	"jwt"
	"sql"
	"github.com/gin-gonic/gin"
)

func HandlerJwt(c *gin.Context) {
	jwtInfoFromC, err := c.Get("jwtInfo")
	if err == false {
		fmt.Println("jwtInfo is not exist")
		c.Abort()
	}
	jwtInfo := jwtInfoFromC.(*jwt.Jwt)
	jwtCookie, _ := c.Cookie("jwtCookie")
	username := c.PostForm("username")
	if len(jwtCookie) > 0 && jwtInfo.CheckToken(jwtCookie) {
		token := jwtInfo.GetSecrectToken(username)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		fmt.Println("token success!", token)
		c.Next()
	} else {

		password := c.PostForm("password")

		if sql.CheckAccount(username, password) {
			token := jwtInfo.GetSecrectToken(username)
			fmt.Println("username and password is correct!", token)
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
			c.Next()
		} else {
			c.String(http.StatusForbidden, "username or password is wrong!")
			c.Abort()
		}
	}
}
