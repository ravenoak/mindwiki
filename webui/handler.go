package webui

import (
	"html/template"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type handler struct {
	t   *template.Template
	d   *SiteData
}

type crudHandler struct {
	handler
	rtr *mux.Router
	t   map[string]*template.Template
}

func (h *handler) funcMap(funcs map[string]interface{}) {
	var m = template.FuncMap{
		"test": func(s1, s2 string) string {
			return s1 + s2
		},
	}

	for k, v := range m {
		if _, ok := funcs[k]; !ok {
			log.Debug().Str("func", k).Msg("adding template func")
			funcs[k] = v
		} else {
			log.Error().Str("func", k).Msg("template func already defined")
		}
	}
}
