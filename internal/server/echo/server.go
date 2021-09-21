package echo

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ravenoak/mindwiki/internal/app"
	"github.com/ravenoak/mindwiki/internal/config"
	"github.com/ravenoak/mindwiki/internal/server/handlers"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 10 * time.Second
)

type (
	server struct {
		config  *config.WebConfig
		echo    *echo.Echo
		storage app.Persistenator
	}
)

func HTTPServer(c *config.AppConfig, st app.Persistenator) (app.HTTPServinator, error) {
	e := echo.New()
	e.Debug = c.DebugMode
	if c.DebugMode {
		e.HideBanner = false
		e.HidePort = false
	} else {
		e.HideBanner = true
		e.HidePort = true
	}
	e.Server.ReadTimeout = readTimeout
	e.Server.WriteTimeout = writeTimeout
	e.Validator = &customValidator{}

	if err := e.Validator.Validate(c); err != nil {
		return nil, err
	}

	s := &server{
		config:  c.WebConfig,
		echo:    e,
		storage: st,
	}

	s.setupRoutes()
	s.setupMiddleware()

	return s, nil
}

func (s *server) Start() error {
	return s.echo.Start(fmt.Sprintf("%s:%d", s.config.Bind, s.config.Port))
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func (s *server) setupMiddleware() {
	s.echo.Use(middleware.Gzip())
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.Secure())
}

func (s *server) setupRoutes() {
	pg := handlers.NewCRUDHandler("pages", s.storage)
	pg.Register(s.echo, "/")
}
