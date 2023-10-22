package main

import (
	"go-framework/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	//handler with basic routing
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
			"version": "not set",
		})
	})

	router.Run(config.APP_PORT)
}
