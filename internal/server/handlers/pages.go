package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PageGroup struct {
	
}

// e.GET("/pages/:id", getPage)
func GetPage(c echo.Context) error {
	// Page ID from path `pages/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
