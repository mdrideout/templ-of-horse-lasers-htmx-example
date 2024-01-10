package leaderboard

import (
	"log"
	"main/app/leaderboard/components"
	"main/app/leaderboard/models"
	"main/db"
	"main/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LeaderboardService(c echo.Context) error {
	template.NewTemplateRenderer(c.Echo())

	// Get params
	title := c.QueryParam("title")

	// Get the updated list of items from the database
	items, err := getLeaderboardTableData()
	if err != nil {
		log.Println(err)
		return err
	}

	// Get the component
	component := Leaderboard(title, items)

	// Render the component
	return template.AssertRender(c, http.StatusOK, component)
}

func LeaderboardAddService(c echo.Context) error {
	template.NewTemplateRenderer(c.Echo())

	// Construct the model and bind the form data to it
	leaderboardItem := &models.LeaderboardItemStore{}
	if err := c.Bind(leaderboardItem); err != nil {
		return err
	}

	// Save to database using GORM
	createResult := db.DB.Create(&leaderboardItem)
	if createResult.Error != nil {
		return createResult.Error
	}

	// Get the updated list of items from the database
	items, err := getLeaderboardTableData()
	if err != nil {
		log.Println(err)
		return err
	}

	// Create the component with the leaderboard data
	component := components.LeaderboardTable(items)

	// Render the component
	return template.AssertRender(c, http.StatusOK, component)
}
