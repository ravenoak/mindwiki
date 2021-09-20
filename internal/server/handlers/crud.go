package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/ravenoak/mindwiki/internal/app"
	"github.com/ravenoak/mindwiki/internal/nodes"
	"github.com/ravenoak/mindwiki/internal/primitives"
	"github.com/rs/zerolog/log"
)

type CRUDHandler struct {
	objectType string
	storage    app.Storinator
	echo       *echo.Group
}

func (h *CRUDHandler) Create(c echo.Context) error {
	return nil
}

func (h *CRUDHandler) Read(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 16)
	a := strings.Split(c.Request().Header["Accept"][0], ",")
	log.Debug().Interface("header-accept", a).Msg("")
	// o, _ := h.storage.Get(id, h.objectType)
	o := nodes.Page{
		Node:  primitives.Node{uint64(id)},
		Title: "Foobar and the Spammy Eggs",
		Body:  "This is the body",
		Slug:  "foobar",
		Tags:  nil,
	}
	for _, v := range a {
		if v == "text/html" {
			return c.String(http.StatusOK, fmt.Sprintf("%#v", o))
		} else if v == "application/json" {
			return c.JSON(http.StatusOK, o)
		} else if strings.Contains(v, "*/*") {
			return c.JSONPretty(http.StatusOK, o, "    ")
		}
	}
	return c.NoContent(http.StatusBadRequest)
}

func (h *CRUDHandler) Update(c echo.Context) error {
	return nil
}

func (h *CRUDHandler) Delete(c echo.Context) error {
	return nil
}

func (h *CRUDHandler) List(c echo.Context) error {
	return nil
}

func (h *CRUDHandler) Find(c echo.Context) error {
	return nil
}

func (h *CRUDHandler) Register(e *echo.Echo, p string, m ...echo.MiddlewareFunc) {
	if p == "" {
		p = "/"
	}
	h.echo = e.Group(p+h.objectType, m...)

	// list
	h.echo.GET("/", h.List)
	// h.echo.GET("/index", h.List)
	// h.echo.GET("/list", h.List)

	// create
	h.echo.POST("/", h.Create)

	// read
	h.echo.GET("/:id", h.Read)

	// update
	h.echo.PUT("/:id", h.Update)
	h.echo.PATCH("/:id", h.Update)

	// delete
	h.echo.DELETE("/:id", h.Delete)
}

func NewCRUDHandler(t string, s app.Storinator) *CRUDHandler {
	return &CRUDHandler{
		objectType: t,
		storage:    s,
	}
}

func makeResponse() {

}
