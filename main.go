package main

import (
	db "web_gin/database"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()  //哈哈哈哈哈
	router.Run(":8080")
}
