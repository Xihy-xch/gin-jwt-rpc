package main

import (
	"fmt"
	"sql"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	sql.Connection("root", "root", "gin_jwt_rpc")
	fmt.Println(sql.CheckAccount("xiht", "root"))
	fmt.Println(sql.CheckAccount("xihy", "root"))
}
