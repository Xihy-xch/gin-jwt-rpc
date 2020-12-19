package main

import (
	"fmt"
	"workspace/gin-jwt-rpc/sql"
)

func main() {
	sql.Connection("root", "root", "gin_jwt_rpc")
	fmt.Println(sql.CheckAccount("xiht", "root"))
	fmt.Println(sql.CheckAccount("xihy", "root"))
}
