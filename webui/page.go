package webui

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"github.com/ravenoak/mindwiki/storage"
)

type pageHandler crudHandler

type pageDisplayData struct {
	SitePage
	P *storage.Page
}

type pageListData struct {
	SitePage
	P []*storage.Page
}

func (h *pageHandler) Display(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Info().Interface("vars", vars).Msg("")
	p := &storage.Page{
		Title: "FooBar Express",
		Body:  "Our meaningless control for mind is to follow others oddly.",
		Slug:  vars["slug"],
	}
	d := &pageDisplayData{
		P: p,
	}
	d.PageTitle = "WikiPage: " + p.Title
	d.SiteData = h.d
	w.WriteHeader(http.StatusOK)
	err := h.t["detail"].ExecuteTemplate(w, "page_detail.gohtml", d)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
}

func (h *pageHandler) List(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Info().Interface("vars", vars).Msg("")

	d := &pageListData{

	}
	d.SiteData = h.d

	d.P = append(d.P, &storage.Page{
		Title: "FooBar Express",
		Body:  "Our meaningless control for mind is to follow others oddly.",
		Slug:  "something",
	})
	w.WriteHeader(http.StatusOK)
	err := h.t["list"].ExecuteTemplate(w, "page_list.gohtml", d)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
}

func (h *pageHandler) Setup() {
	//log.Debug().Msg(h.t.DefinedTemplates())
	h.rtr.HandleFunc("/", h.List).Methods("GET")
	h.rtr.HandleFunc("/{slug}", h.Display).Methods("GET")
	h.rtr.HandleFunc("/{slug}/", h.Display).Methods("GET")
}
