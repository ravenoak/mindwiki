package app

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/ravenoak/mindwiki/storage"
	"github.com/ravenoak/mindwiki/webui"
)

const (
	errConfigNil      = "ummm, `App.Config` can't be nil my dude"
	errWebUIServerNil = "`App.webUIServer` is nil and it's your fault"
)

type (
	App struct {
		Config      *Config
		webUIServer *webui.WebUIServer
		storage     *storage.Storage
	}
)

func (a *App) StartWebUI() error {
	if a.Config == nil {
		return errors.New(errConfigNil)
	}
	if a.webUIServer == nil {
		addr := fmt.Sprintf("%s:%d", a.Config.WebUIBind, a.Config.WebUIPort)
		a.webUIServer = webui.NewServer(addr)
	}
	return a.webUIServer.Start()
}

func (a *App) StopWebUI() error {
	if a.webUIServer == nil {
		log.Fatal().Err(errors.New(errWebUIServerNil))
	}
	return a.webUIServer.Stop()
}
