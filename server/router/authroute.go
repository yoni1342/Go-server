package router

import (
	"server/controller"

	"github.com/gin-gonic/gin"
)

func GetAuthGroup(route *gin.Engine) *gin.RouterGroup {
	return route.Group("/auth")
}

func AuthRoutes(route *gin.RouterGroup) {
	route.POST("/signup", controller.Signup)
	route.POST("/signin", controller.Signin)
	route.POST("/verifyemail", controller.VerifyEmail)
}
