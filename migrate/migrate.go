package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

// マイグレーションを実行するための関数
func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Success!! Migration")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
