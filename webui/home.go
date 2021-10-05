package webui

import (
	"html/template"
	"net/http"

	"github.com/rs/zerolog/log"
)

type homeHandler struct {
	t *template.Template
}

type homeData struct {
	Name      string
	PageTitle string
}

func (h *homeHandler) Display(w http.ResponseWriter, r *http.Request) {
	d := homeData{
		Name:      "World",
		PageTitle: "Home",
	}
	w.WriteHeader(http.StatusOK)
	err := h.t.ExecuteTemplate(w, "home.gohtml", d)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
}
