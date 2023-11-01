package router

import (
	"server/controller"

	"github.com/gin-gonic/gin"
)

func GetUploadGroup(route *gin.Engine) *gin.RouterGroup {
	return route.Group("/upload")
}

func UploadRoutes(route *gin.RouterGroup) {
	route.POST("/", controller.UploadFile)
}
