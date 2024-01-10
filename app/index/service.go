package index

import (
	"main/app/index/components"
	"main/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexRoute(c echo.Context) error {
	template.NewTemplateRenderer(c.Echo())

	// Get params
	name := c.QueryParam("name")

	// Get the component
	component := Index(name)

	// Render the component
	return template.AssertRender(c, http.StatusOK, component)
}

func IndexRouteYourNameIs(c echo.Context) error {
	template.NewTemplateRenderer(c.Echo())

	// Get params
	name := c.QueryParam("name")

	// Get the component
	component := components.YourNameIs(name)

	// Render the component
	return template.AssertRender(c, http.StatusOK, component)
}

func IndexRouteYourNameIsGlobal(c echo.Context) error {
	template.NewTemplateRenderer(c.Echo())

	// Get params
	newGlobalName := c.FormValue("name")

	// Update the global variable
	components.GlobalName = newGlobalName

	// Get the component
	component := components.YourNameIsGlobal()

	// Render the component
	return template.AssertRender(c, http.StatusOK, component)
}
