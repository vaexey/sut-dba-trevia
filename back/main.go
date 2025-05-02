package main

import (
	"back/auth"
	"back/config"
	"back/db"
	"back/routes"
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
		config.Config.Database.Port,
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

	authHandler := auth.Handler{
		Db: &database,
	}

	authMiddleware := authHandler.RequireJWT()

	routes := routes.NewApi(&database)

	api := router.Group(config.Config.ApiPath)
	{
		api.POST("/login", authHandler.Login)
		api.POST("/sign-up", authHandler.Register)

		api.Use(authMiddleware) 
		{
			api.GET("/user", routes.GetCurrentUser)
		}
	}

	router.NoRoute(func(ctx *gin.Context){
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "Not found"})
	})

	addr := ":" + strconv.Itoa(config.Config.Server.Port)
	router.Run(addr)
}

