package webui

import (
	"context"
	"errors"
	"html/template"
	"net/http"
	"time"

	"github.com/benbjohnson/hashfs"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

const (
	stopTimeout  = 15 * time.Second
	idleTimeout  = 60 * time.Second
	readTimeout  = 15 * time.Second
	writeTimeout = 15 * time.Second

	errSvrNil = "awww, you've gone and done it now; `s.svr` is nil"
	errRtrNil = "dude, where's your router?; `s.rtr` is nil"
)

type WebUIServer struct {
	addr string

	rtr *mux.Router
	svr *http.Server

	static http.Handler

	d  *SiteData
	hh *homeHandler
	ph *pageHandler
}

func (s *WebUIServer) Start() error {
	if s.svr == nil {
		return errors.New(errSvrNil)
	}
	log.Info().Str("WebUIServer.addr", s.addr).Msg("starting webui")
	err := s.svr.ListenAndServe()
	if err != nil {
		log.Error().Err(err)
	}
	return err
}

func (s *WebUIServer) Stop() error {
	if s.svr == nil {
		return errors.New(errSvrNil)
	}
	log.Info().Interface("stopTimeout", stopTimeout).Msg("stopping webui")
	ctx, cancel := context.WithTimeout(context.Background(), stopTimeout)
	defer cancel()
	err := s.svr.Shutdown(ctx)
	if err != nil {
		log.Error().Err(err)
	}
	return err
}

func (s *WebUIServer) setupHome() {
	if s.rtr == nil {
		log.Fatal().Err(errors.New(errRtrNil))
	}

	s.hh = &homeHandler{
		t: template.Must(template.ParseFS(templateFS, "templates/home.gohtml", "templates/layout/*.gohtml")),
		d: s.d,
	}

	s.rtr.HandleFunc("/", s.hh.Display).Methods("GET")
}

func (s *WebUIServer) setupPages() {
	if s.rtr == nil {
		log.Fatal().Err(errors.New(errRtrNil))
	}

	s.ph = &pageHandler{
		rtr: s.rtr.PathPrefix("/page").Subrouter(),
		t: make(map[string]*template.Template),
	}

	s.ph.t["detail"] = template.Must(template.ParseFS(
		templateFS,
		"templates/pages/page_detail.gohtml",
		"templates/layout/*.gohtml",
	))
	s.ph.t["list"] = template.Must(template.ParseFS(
		templateFS,
		"templates/pages/page_list.gohtml",
		"templates/layout/*.gohtml",
	))
	s.ph.d = s.d
	s.ph.Setup()
}

func NewServer(addr string) *WebUIServer {
	r := mux.NewRouter()
	st := hashfs.FileServer(staticFS)

	r.Use(loggingMiddleware)
	// r.Use(handlers.RecoveryHandler())

	s := &WebUIServer{
		addr: addr,
		rtr:  r,
		svr: &http.Server{
			Addr:         addr,
			Handler:      r,
			IdleTimeout:  idleTimeout,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		},
		d: &SiteData{
			Links: &SiteLinks{
				Index:     "/",
				PageIndex: "/page/",
			},
		},
		static: st,
	}

	r.Handle("/static", st)
	s.setupHome()
	s.setupPages()

	return s
}
