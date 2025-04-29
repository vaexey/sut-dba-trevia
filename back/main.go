package main

import (
	"back/auth"
	"back/config"
	"back/db"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", 
		config.Config.Database.Host,
		5432,
		config.Config.Database.Username,
		config.Config.Database.Database,
		config.Config.Database.Password,
	)

	dbConn, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
    	panic("failed to create database connection: " + err.Error())
	}

	sqlDB, err := dbConn.DB()
	if err != nil {
    	panic("failed to get database instance: " + err.Error())
	}

	err = sqlDB.Ping()
	if err != nil {
    	panic("failed to ping database: " + err.Error())
	}

	database := db.NewDatabase(dbConn)

	// auto migration
	if err = database.Migrate(); err != nil {
		panic("failed to create a database" + err.Error())
	}

	api := router.Group(config.Config.ApiPath)
	{
		api.GET("/hello", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{"msg": "Hello world"})
		})
		
		api.POST("/login", auth.HandleLogin)
		api.GET("/config", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Host": config.Config.Database.Host,
				"Port": config.Config.Database.Port,
				"Username": config.Config.Database.Username,
				"Password": config.Config.Database.Password,
				"Name": config.Config.Database.Database,
			})
		})
		api.Use(auth.AuthMiddleware()) 
		{
			// TODO: auth endpoints 
		}
	}

	router.NoRoute(func(ctx *gin.Context){
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "Not found"})
	})

	addr := ":" + strconv.Itoa(config.Config.Server.Port)
	router.Run(addr)
}

