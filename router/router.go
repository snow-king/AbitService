package router

import (
	PersonControllers "AbitService/app/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/person/:id", PersonControllers.Index)
	r.GET("/person/:id/family", PersonControllers.ShowFamily)
	return r
}
