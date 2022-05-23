package account

import (
	"github.com/Etwodev/Doctorate/server/router"
)

func NewRouter(status bool) router.Router {
	return router.NewRouter(initRoutes(), status)
}

func initRoutes() []router.Route {
	return []router.Route{
		router.NewPostRoute("/ABCDEFGHIJKLMNO/account/yostar_auth_request", true, true, EmailVerificationPostRoute),
		// router.NewPostRoute("/ABCDEFGHIJKLMNO/account/yostar_auth_submit", true, true, EmailSubmitPostRoute),
	}
}