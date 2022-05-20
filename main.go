package main

import (

	// SYSTEM
	"log"
	"net/http"
	// EXTERNAL
	"github.com/joho/godotenv"
	// INTERNAL
	"github.com/Etwodev/Doctorate/mitm"
	"github.com/Etwodev/Doctorate/server"
	"github.com/Etwodev/Doctorate/server/router"

	// ROUTERS
	"github.com/Etwodev/Doctorate/server/router/assetbundle"
	"github.com/Etwodev/Doctorate/server/router/config"
)

func main () {
	err := godotenv.Load()

	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	var c = server.CreateConfig()
	var s = server.New(c)

	routers := []router.Router{
		config.NewRouter(true),
		assetbundle.NewRouter(true),
	}

	s.InitRouter(routers...)

	go func() {
		mitm.TestingMain("1", "tcp", "127.0.0.1", "4000")
	}()

	var handler = s.InitRouters()

	log.Fatal(http.ListenAndServe(c.Doctorate.Port, handler))
}






