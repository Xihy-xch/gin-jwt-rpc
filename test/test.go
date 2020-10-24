package main

import (
	"base64"
	"fmt"
	"jwt"
	"strings"
)

type A struct {
	Usr string
}
type User struct {
	Id   int
	Name string
	//addr string
}

func main() {

	jwtInfo := jwt.NewJwt()
	jwtInfo.SetHeader("abc")
	// jwtInfo.SetPlayload("xiaoming")
	jwtInfo.SetSecret("secret")

	fmt.Println(jwtInfo)
	fmt.Println(jwtInfo.GetHeaderToBase64())
	fmt.Println(jwtInfo.GetPayloadToBase64())
	var str strings.Builder
	str.WriteString(jwtInfo.GetHeaderToBase64())
	str.WriteString(".")
	str.WriteString(jwtInfo.GetPayloadToBase64())
	fmt.Println(base64.DecodeFromBase64(jwtInfo.GetHeaderToBase64()))
	fmt.Println(base64.DecodeFromBase64(jwtInfo.GetPayloadToBase64()))
	fmt.Println(jwtInfo.ComputeHmacSha256(str.String()))
}
