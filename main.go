package main

import (
	"os"
	"github.com/Etwodev/Doctorate/server"
	"github.com/Etwodev/Doctorate/server/router"
	"github.com/Etwodev/Doctorate/server/router/account"
	"github.com/Etwodev/Doctorate/server/router/announce"
	"github.com/Etwodev/Doctorate/server/router/app"
	"github.com/Etwodev/Doctorate/server/router/assetbundle"
	"github.com/Etwodev/Doctorate/server/router/config"
	"github.com/Etwodev/Doctorate/server/router/user"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	connection := os.Getenv("HOST_ROUTER_DSM")
	name := os.Getenv("HOST_ROUTER_NAME")
	address := os.Getenv("HOST_ROUTER_ADDRESS")
	version := os.Getenv("HOST_ROUTER_VERSION")
	port := os.Getenv("HOST_ROUTER_PORT")
	experimental := (os.Getenv("HOST_ROUTER_EXPERIMENTAL") == "true")

	s := server.Server{
		Version: version,
		Port: port,
		Experimental: experimental,
		Name: name,
		Address: address,
		Connection: connection,
	}

	s.Initilise()

	routers := []router.Router{
		config.NewRouter(true),
		assetbundle.NewRouter(true),
		account.NewRouter(true),
		app.NewRouter(true),
		user.NewRouter(true),
		announce.NewRouter(true),
	}

	s.Start(routers...)
}
