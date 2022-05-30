package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Etwodev/Doctorate/server/helpers"
	"github.com/Etwodev/Doctorate/server/router"
	stat "github.com/Etwodev/Doctorate/static"

	this "log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type RouterConfig struct {
	Version			string
	Port			string
	Experimental	bool
	Name			string
	Address			string
}

type DatabaseConfig struct {
	Address			string
	Username		string
	Password		string
	Target			string
	Format			string
	Connection		string
}

type Config struct {
	Router		*RouterConfig
	SQL  		*DatabaseConfig
}

type Server struct {
	cfg       *Config
	routers   []router.Router
}

func New() *Server {
	return &Server{
		cfg: initConf(),
	}
}

func (s *Server) Start(routers ...router.Router)  {
	s.routers = append(s.routers, routers...)
	err := helpers.Updater(stat.IOS)
	if err != nil {
		log.Debug().Msgf("Start: failed to get hotupdate data: %s", err)
	}
	err = helpers.Updater(stat.Android)
	if err != nil {
		log.Debug().Msgf("Start: failed to get hotupdate data: %s", err)
	}
	h := s.handler()
	this.Fatal(http.ListenAndServe(s.cfg.Router.Port, h))
}

func (s *Server) handler() *chi.Mux {
	m := s.createMux()
	return m
}

func (s *Server) createMux() *chi.Mux {
	m := chi.NewMux()
	m.Use(middleware.RequestID)
	m.Use(middleware.RealIP)
	m.Use(middleware.Logger)
	m.Use(middleware.Recoverer)
	for _, router := range s.routers {
		if ( router.Status() ) {
			for _, r := range router.Routes() {
				if ( r.Status() && ( r.Experimental() == s.cfg.Router.Experimental || !r.Experimental() ) ) {
					log.Info().Bool("Experimental", r.Experimental()).Bool("Status", r.Status()).Str("Method", r.Method()).Str("Path", r.Path()).Msg("Registering route")
					m.Method(r.Method(), r.Path(), r.Handler())
				}
			}
		}
	}
	return m
}

func initConf() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	form := os.Getenv("HOST_SQL_FORMAT")
	user := os.Getenv("HOST_SQL_USER")
	pass := os.Getenv("HOST_SQL_PASS")
	addr := os.Getenv("HOST_SQL_ADDRESS")
	tabl := os.Getenv("HOST_SQL_TABLE")
	conn := fmt.Sprintf(form, user, pass, addr, tabl)
	helpers.Connect(conn)

	db := &DatabaseConfig{
		Address: addr,
		Username: user,
		Password: pass,
		Target: tabl,
		Format: form,
		Connection: conn,
	}

	nom := os.Getenv("HOST_ROUTER_NAME")
	adr := os.Getenv("HOST_ROUTER_ADDRESS")
	ver := os.Getenv("HOST_ROUTER_VERSION")
	prt := os.Getenv("HOST_ROUTER_PORT")
	exp := (os.Getenv("HOST_ROUTER_EXPERIMENTAL") == "true")

	rc := &RouterConfig{
		Version: ver,
		Port: prt,
		Experimental: exp,
		Name: nom,
		Address: adr,
	}

	return &Config{
		Router: rc,
		SQL: db,
	}
}
