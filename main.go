package main

import (
	db "web_gin/database"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8080")
}