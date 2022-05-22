package server

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Etwodev/Doctorate/server/helpers"
	"github.com/Etwodev/Doctorate/server/router"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

type DoctorateConfig struct {
	Version			string
	Port			string
	Experimental	bool
	Downloads		string
}

type SQLConfig struct {
	Address			string
	Username		string
	Password		string
	Target			string
	Format			string
	Connection		string
}

type Config struct {
	Doctorate		*DoctorateConfig
	SQL  			*SQLConfig
}

type Server struct {
	cfg       *Config
	routers   []router.Router
}

func New(cfg *Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) InitRouter(routers ...router.Router) {
	s.routers = append(s.routers, routers...)
}

func (s *Server) InitRouters() *chi.Mux {
	return s.createMux()
}

func CreateConfig() *Config {
	dc := getDC()
	sc := getSC()

	c := &Config{
		Doctorate: dc,
		SQL: sc,
	}

	return c
}

func getDC() *DoctorateConfig {
	b, err := strconv.ParseBool(os.Getenv("HOST_DOCTORATE_EXPERIMENTAL"))
	if err != nil {
		panic(err)
	}

	dc := &DoctorateConfig{
		Version: os.Getenv("HOST_DOCTORATE_VERSION"),
		Port: os.Getenv("HOST_DOCTORATE_PORT"),
		Experimental: b,
		Downloads: os.Getenv("HOST_DOWNLOAD_URL"),
	}
	return dc
}

func getSC() *SQLConfig {
	fm := os.Getenv("HOST_SQL_FORMAT")
	us := os.Getenv("HOST_SQL_USER")
	ps := os.Getenv("HOST_SQL_PASS")
	ad := os.Getenv("HOST_SQL_SERVER")
	db := os.Getenv("HOST_SQL_DB")
	cn := fmt.Sprintf(fm, us, ps, ad, db)

	sc := &SQLConfig{
		Address: ad,
		Username: us,
		Password: ps,
		Target: db,
		Format: fm,
		Connection: cn,
	}

	helpers.Connect(cn)

	return sc
}

func (s *Server) createMux() *chi.Mux {
	m := chi.NewMux()

	m.Use(middleware.RequestID)
	m.Use(middleware.RealIP)
	m.Use(middleware.Logger)
	m.Use(middleware.Recoverer)
	
	helpers.HotUpdater(s.cfg.Doctorate.Downloads, "IOS")
	helpers.HotUpdater(s.cfg.Doctorate.Downloads, "Android")

	log.Debug().Msg("Registering routers")
	for _, router := range s.routers {
		if ( router.Status() ) {
			for _, r := range router.Routes() {
				if ( r.Status() && ( r.Experimental() == s.cfg.Doctorate.Experimental || !r.Experimental() ) ) {
					log.Debug().Bool("Experimental", r.Experimental()).Bool("Status", r.Status()).Str("Method", r.Method()).Str("Path", r.Path()).Msg("Registering route")
					m.Method(r.Method(), r.Path(), r.Handler())
				}
			}
		}
	}
	log.Debug().Msg("Register complete")
	return m
}