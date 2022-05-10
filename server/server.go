package server

import (
	"github.com/Etwodev/Doctorate/server/router"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Version  	string
	Port        string
	Connection  string
}

type Server struct {
	cfg       *Config
	routers   []router.Router
}


func New(cfg *Config) *Server {
	router.Connect(cfg.Connection)
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) InitRouter(routers ...router.Router) {
	s.routers = append(s.routers, routers...)
}

func (s *Server) InitRouters(experimental bool) *chi.Mux {
	return s.createMux(experimental)
}

func (s *Server) createMux(experimental bool) *chi.Mux {
	m := chi.NewMux()

	m.Use(middleware.RequestID)
	m.Use(middleware.RealIP)
	m.Use(middleware.Logger)
	m.Use(middleware.Recoverer)

	log.Debug().Msg("Registering routers")
	for _, apiRouters := range s.routers {
		if ( apiRouters.Status() ) {
			for _, r := range apiRouters.Routes() {
				if ( r.Status() && ( r.Experimental() == experimental || !r.Experimental() ) ) {
					log.Debug().Bool("Experimental", r.Experimental()).Bool("Status", r.Status()).Str("Method", r.Method()).Str("Path", r.Path()).Msg("Registering route")
					m.Method(r.Method(), r.Path(), r.Handler())
				}
			}
		}
	}
	log.Debug().Msg("Register complete")
	return m
}