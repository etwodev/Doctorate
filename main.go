package main

import (
	// INTERNAL
	"github.com/Etwodev/Doctorate/server"
	"github.com/Etwodev/Doctorate/server/router"

	// ROUTERS
	"github.com/Etwodev/Doctorate/server/router/account"
	"github.com/Etwodev/Doctorate/server/router/app"
	"github.com/Etwodev/Doctorate/server/router/assetbundle"
	"github.com/Etwodev/Doctorate/server/router/config"
	"github.com/Etwodev/Doctorate/server/router/user"
)

func main () {
	var s = server.New()

	routers := []router.Router{
		config.NewRouter(true),
		assetbundle.NewRouter(true),
		account.NewRouter(true),
		app.NewRouter(true),
		user.NewRouter(true),
	}

	s.Start(routers...)
}