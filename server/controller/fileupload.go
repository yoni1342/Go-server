package controller

import (
	"server/types"
	"server/utils/file"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {
	var inputImages = types.UploadInput{}

	if err := ctx.ShouldBindJSON(&inputImages); err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid request",
			"err":     err,
		})
	}

	var urls []string
	images := inputImages.Input.Args

	for _, image := range images {
		url, err := file.UploadFile(image.Base64Data)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "invalid request",
			})
		}
		urls = append(urls, url)
	}

	ctx.JSON(200, gin.H{
		"urls": urls,
	})
}
