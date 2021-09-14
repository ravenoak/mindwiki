package server

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ravenoak/mindwiki/internal/app"
	"github.com/ravenoak/mindwiki/internal/config"
	"github.com/ravenoak/mindwiki/internal/server/handlers"
	"github.com/ravenoak/mindwiki/internal/storage"
	"github.com/rs/zerolog/log"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 10 * time.Second
)

type (
	server struct {
		config      *config.WebConfig
		httpServer  *echo.Echo
		storage     *storage.Depot
	}
)

func New(c *config.AppConfig, st *storage.Depot) (app.HTTPServer, error) {
	e := echo.New()
	e.Debug = c.DebugMode
	e.HideBanner = false
	e.HidePort = false
	e.Server.ReadTimeout = readTimeout
	e.Server.WriteTimeout = writeTimeout
	e.Validator = &customValidator{}

	if err := e.Validator.Validate(c); err != nil {
		return nil, err
	}

	s := &server{
		config:     c.WebConfig,
		httpServer: e,
		storage: st,
	}

	s.setupRoutes()
	s.setupMiddleware()

	return s, nil
}

func (s server) Start() error {
	log.Debug().Msg(fmt.Sprintf("Port: %d", *s.config.Port))
	return s.httpServer.Start(fmt.Sprintf("%s:%d", s.config.Bind, *s.config.Port))
}

func (s server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s server) setupMiddleware() {
	s.httpServer.Use(middleware.Gzip())
	s.httpServer.Use(middleware.Logger())
	s.httpServer.Use(middleware.Recover())
	s.httpServer.Use(middleware.Secure())
}

func (s server) setupRoutes() {
	s.httpServer.GET("/pages/:id", handlers.GetPage)
}
