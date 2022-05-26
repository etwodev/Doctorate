package user

import (
	"github.com/Etwodev/Doctorate/server/helpers"
	"github.com/Etwodev/Doctorate/server/router"
)

func NewRouter(status bool) router.Router {
	helpers.Engine.CreateTables(&AuthData{})
	return router.NewRouter(initRoutes(), status)
}

func initRoutes() []router.Route {
	return []router.Route{
		router.NewPostRoute("/ABCDEFGHIJKLMNO/user/agreement", true, true, AgreementPostRoute),
		router.NewPostRoute("/ABCDEFGHIJKLMNO/user/create", true, true, CreatePostRoute),
	}
}

