package server

import (
	"net/http"
	"github.com/Etwodev/Doctorate/server/helpers"
	"github.com/Etwodev/Doctorate/server/router"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
	this "log"
)

type Server struct {
	Version			string
	Port			string
	Experimental	bool
	Name			string
	Address			string
	Connection		string
	routers   		[]router.Router
}

func (s *Server) Start(routers ...router.Router)  {	
	s.routers = append(s.routers, routers...)
	this.Fatal(http.ListenAndServe(s.Port, s.handler()))
}

func (s *Server) Initilise() {
	err := helpers.Connect(s.Connection)
	if err != nil {
		log.Debug().Msgf("Start: failed to connect to SQL: %s", err)
	}

	err = helpers.Updater("IOS")
	if err != nil {
		log.Debug().Msgf("Start: failed to get hotupdate data: %s", err)
	}

	err = helpers.Updater("Android")
	if err != nil {
		log.Debug().Msgf("Start: failed to get hotupdate data: %s", err)
	}

	err = helpers.Init()
	if err != nil {
		log.Debug().Msgf("Start: failed to initilise static variables: %s", err)
	}
}

func (s *Server) handler() *chi.Mux {
	m := chi.NewMux()
	// If we need to handle more middleware, seperate
	m.Use(middleware.RequestID)
	m.Use(middleware.RealIP)
	m.Use(middleware.Logger)
	m.Use(middleware.Recoverer)
	s.initMux(m)
	return m
}

func (s *Server) initMux(m *chi.Mux) {
	for _, router := range s.routers {
		if ( router.Status() ) {
			for _, r := range router.Routes() {
				if ( r.Status() && ( r.Experimental() == s.Experimental || !r.Experimental() ) ) {
					log.Info().Bool("Experimental", r.Experimental()).Bool("Status", r.Status()).Str("Method", r.Method()).Str("Path", r.Path()).Msg("Registering route")
					m.Method(r.Method(), r.Path(), r.Handler())
				}
			}
		}
	}
}
