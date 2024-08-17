package main

import (
	"github.com/gin-gonic/gin"
	"judoApi/controllers"
)

func InitializeRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Judoka endpoint
	router.GET("/judokas", controllers.GetJudokas)
	router.GET("/judoka/:id", controllers.GetJudoka)
	router.POST("/judoka", controllers.CreateJudoka)

	// Club endpoint
	router.GET("/clubs", controllers.GetClubs)
	router.GET("/club/:id", controllers.GetClub)
	router.POST("/club", controllers.CreateClub)

	// Tournament endpoint
	router.GET("/tournaments", controllers.GetTournaments)
	router.GET("/tournament/:id", controllers.GetTournament)
	router.POST("/tournament", controllers.CreateTournament)

	// Categorie endpoint
	router.GET("/categories/:tournamentID", controllers.GetCategories)
	router.GET("/categories/:tournamentID/open", controllers.GetOpenCategories)
	router.GET("/category/:id", controllers.GetCategory)
	router.POST("/category", controllers.CreateCategory)

	// Categorie endpoint
	router.POST("/inscription", controllers.CreateRegistration)
	router.GET("/inscriptions", controllers.GetRegistrationsPerTournament)

	return router
}
