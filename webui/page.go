package webui

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"github.com/ravenoak/mindwiki/storage"
)

type pageHandler struct {
	rtr *mux.Router
	t *template.Template
}

type pageData struct {
	P         *storage.Page
	PageTitle string
	Name string
}

func (h *pageHandler) Display(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Info().Interface("vars", vars).Msg("")
	p := &storage.Page{
		Title: "FooBar Express",
		Body: "Our meaningless control for mind is to follow others oddly.",
		Slug: vars["slug"],
	}
	d := &pageData{
		P:         p,
		PageTitle: "WikiPage: " + p.Title,
		Name:      "Fuck me... :-(",
	}
	w.WriteHeader(http.StatusOK)
	err := h.t.ExecuteTemplate(w, "page_detail.gohtml", d)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
}

func (h *pageHandler) Setup() {
	h.rtr.HandleFunc("/{slug}", h.Display).Methods("GET")
	h.rtr.HandleFunc("/{slug}/", h.Display).Methods("GET")
}
