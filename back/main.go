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
		// user
		api.POST("/login", authHandler.Login)
		api.POST("/sign-up", authHandler.Register)

		// loaction
		api.GET("/locations/:locationId", routes.LocationsById)
		api.GET("/locations/search", routes.LocationSearch)

		// attractions
		api.GET("/attractions/:attractionId", routes.AttractionById)
		api.GET("/attractions/location/:locationId", routes.AttractionByLocation)
		api.GET("/attractions/funfact", routes.AttractionWithRandomFunFact)


		api.Use(authMiddleware)
		{
			// user
			api.GET("/user", routes.GetCurrentUser)

			// attractions
			api.POST("/attractions", routes.CreateAttraction)

			// rating
			api.POST("/rate", routes.CreateRating)

			// comment
			api.POST("/comments", routes.CreateComment)
			api.GET("/comments/:attractionId", routes.GetComments)

			// reports
			api.POST("/reports/attractions", routes.CreateAttractionReport)
			api.POST("/reports/comments", routes.CreateCommentReport)
			api.GET("/reports/attractions", authHandler.RequireModerator(), routes.AttractionReports)
			api.GET("/reports/comments", authHandler.RequireModerator(), routes.CommentReports)

			// stats 
			api.GET("/stats", authHandler.RequireAdmin(), routes.Stats)
		}
	}

	router.NoRoute(func(ctx *gin.Context){
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "Not found"})
	})

	addr := ":" + strconv.Itoa(config.Config.Server.Port)
	router.Run(addr)
}

