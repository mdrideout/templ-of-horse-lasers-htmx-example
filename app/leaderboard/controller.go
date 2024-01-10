package leaderboard

import (
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/leaderboard", LeaderboardService)
	e.POST("/leaderboard/add", LeaderboardAddService)
}
