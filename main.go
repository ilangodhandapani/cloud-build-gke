package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var APP *gin.Engine

func getHomePageHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "hey it is up from V2!"})
}

func main() {
	APP = gin.Default()
	APP.GET("/status", getHomePageHandler)
	PORT := "8080"
	APP.Run(":" + PORT)
}
