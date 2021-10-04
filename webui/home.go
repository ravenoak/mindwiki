package webui

import (
	"html/template"
	"net/http"
)

type homeHandler struct {
	t *template.Template
}

type homeData struct {
	s SiteData
}

func (h *homeHandler) Display(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	h.t.ExecuteTemplate(w, "home.gohtml", "World")
}

func newHomeHandler() *homeHandler {
	return nil
}
