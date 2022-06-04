package user

import (
	// "github.com/Etwodev/Doctorate/server/helpers"
	"github.com/Etwodev/Doctorate/server/router"
)

func NewRouter(status bool) router.Router {
	// Initilise authentication table
	return router.NewRouter(initRoutes(), status)
}

func initRoutes() []router.Route {
	return []router.Route{
		router.NewPostRoute("/ABCDEFGHIJKLMNO/user/agreement", true, true, AgreementPostRoute),
		// router.NewPostRoute("/user/yostar_createlogin", true, true, CreatePostRoute),
		// router.NewPostRoute("/user/login", true, true, CreatePostRoute),
	}
}
