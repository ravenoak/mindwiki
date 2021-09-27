package webui

import (
	"html/template"
	"net/http"
)

type homeHandler struct {
	displayTemplate *template.Template
}

type homeData struct {

}

func (h *homeHandler) Display(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	h.displayTemplate.Execute(w, "World")
}

func newHomeHandler() *homeHandler {
	return nil
}