package webui

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

type homeHandler handler

type homeData struct {
	SitePage
	Name string
}

func (h *homeHandler) Display(w http.ResponseWriter, r *http.Request) {
	d := homeData{
		Name: "World",
	}
	d.PageTitle = "Home"
	d.SiteData = h.d
	w.WriteHeader(http.StatusOK)
	err := h.t.ExecuteTemplate(w, "home.gohtml", d)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
}
