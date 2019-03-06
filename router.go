package main

import (
	"github.com/gin-gonic/gin"
	. "web_gin/apis"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi) //初始化

	router.POST("/person", AddPersonApi) //添加

	router.GET("/persons", GetPersonsApi) //查询所有

	router.GET("/person/:id", GetPersonApi) //查询单个

	router.PUT("/person/:id", ModPersonApi) //更新1

	router.DELETE("/person/:id", DelPersonApi) //删除

	router.POST("/person/:id", UpPersonApi) //更新2

	return router
}
