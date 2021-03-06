package routers

import (
	"github.com/Blac-Panda/Stardome-API/configurations"
	"github.com/Blac-Panda/Stardome-API/controllers"
	"github.com/Blac-Panda/Stardome-API/middlewares"
	"github.com/Blac-Panda/Stardome-API/repositories/database"
	"github.com/Blac-Panda/Stardome-API/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	authenticationHandler controllers.AuthenticationController
	playerHandler         controllers.PlayerController
)

func init() {

	playerRepository := database.NewPlayerRepository(func() *gorm.DB {
		db, err := configurations.GetDB()

		if err != nil {
			return nil
		}

		return db
	})

	playerService := services.NewPlayerService(playerRepository)
	authenticationService := services.NewAuthenticationService(playerRepository)

	playerHandler = controllers.InitPlayerController(playerService)
	authenticationHandler = controllers.InitAuthenticationController(authenticationService)
}

// Routers This  function defines the available routes
// in Stardome API
func Routers() *gin.Engine {
	g := gin.Default()

	g.Use(middlewares.ErrorHandlerMiddleware())

	api := g.Group("/api")
	{
		api.POST("/auth/token", authenticationHandler.AuthenticatePlayer)
		api.POST("/players", playerHandler.CreatePlayer)

		auth := api.Group("", middlewares.AuthHandlerMiddleware())
		{
			auth.GET("/players", playerHandler.ListPlayers)

			auth.GET("/players/:id", playerHandler.GetPlayer)
			auth.PUT("/players/:id", playerHandler.UpdatePlayer)
			auth.PATCH("/players/:id", playerHandler.ModifyPlayer)
			auth.DELETE("/players/:id", playerHandler.DeletePlayer)

			auth.GET("/tournaments", controllers.ListTournaments)
			auth.POST("/tournaments", controllers.CreateTournament)

			auth.GET("/tournaments/:id", controllers.GetTournament)
			auth.PUT("/tournaments/:id", controllers.UpdateTournament)
			auth.PATCH("/tournaments/:id", controllers.ModifyTournament)
			auth.DELETE("/tournaments/:id", controllers.DeleteTournament)
		}
	}

	return g
}
