package main

import (
	"back/auth"
	"back/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group(config.Config.ApiPath)
	{
		api.GET("/hello", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{"msg": "Hello world"})
		})
		
		api.POST("/login", auth.HandleLogin)

		api.Use(auth.AuthMiddleware()) 
		{
			// TODO: auth endpoints 
		}
	}

	router.NoRoute(func(ctx *gin.Context){
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "not found"})
	})

	addr := ":" + strconv.Itoa(config.Config.Server.Port)
	router.Run(addr)
}

