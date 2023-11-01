package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	authRoutes := GetAuthGroup(router)
	AuthRoutes(authRoutes)

	uploadRoutes := GetUploadGroup(router)
	UploadRoutes(uploadRoutes)
}
