package main

import (
	"server/router"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	router.SetupRoutes(route)

	route.Run(":8082")
}
