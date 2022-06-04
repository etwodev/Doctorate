package app


import (
	"github.com/Etwodev/Doctorate/server/router"
)

func NewRouter(status bool) router.Router {
	return router.NewRouter(initRoutes(), status)
}

func initRoutes() []router.Route {
	return []router.Route{
		router.NewPostRoute("/ABCDEFGHIJKLMNO/app/getSettings", true, true, AppSettingsPostRoute),
		router.NewPostRoute("/ABCDEFGHIJKLMNO/app/getCode", true, true, AppCodesPostRoute),
	}
}
