package index

import (
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", IndexRoute)
	e.GET("/index/your-name-is", IndexRouteYourNameIs)
	e.POST("/index/new-global-name", IndexRouteYourNameIsGlobal)
}
