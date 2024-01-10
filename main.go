package main

import (
	"main/app/index"
	"main/app/leaderboard"
	leaderboardModels "main/app/leaderboard/models"

	"main/db"
	"main/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Connect to the database
	db.Connect()
	db.DB.AutoMigrate(&leaderboardModels.LeaderboardItemStore{})

	// Create echo instance
	e := echo.New()

	// Other Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Configure the static file directory
	e.Static("/dist", "dist")

	// Set up the renderer
	template.NewTemplateRenderer(e)

	// Init routes
	index.InitRoutes(e)
	leaderboard.InitRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
