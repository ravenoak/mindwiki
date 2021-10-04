package webui

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type pageHandler struct {
	rtr *mux.Router
}

func (h *pageHandler) Display(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Info().Interface("vars", vars)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Slug: %v\n", vars["slug"])
}

func (h *pageHandler) Setup() {
	h.rtr.HandleFunc("/{slug}", h.Display).Methods("GET")
	h.rtr.HandleFunc("/{slug}/", h.Display).Methods("GET")
}
