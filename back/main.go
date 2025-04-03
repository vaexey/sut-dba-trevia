package main

import (
	"back/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group(config.Config.ApiPath)
	{
		api.GET("/hello", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{"msg": "Hello world2"})
		})


		api.GET("/config", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, config.Config)
		})
	}

	router.NoRoute(func(ctx *gin.Context){
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "not found"})
	})
	addr := ":" + strconv.Itoa(config.Config.Server.Port)
	router.Run(addr)
}