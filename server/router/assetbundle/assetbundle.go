package assetbundle

import (
	"github.com/Etwodev/Doctorate/server/router"
)

func NewRouter(status bool) router.Router {
	return router.NewRouter(initRoutes(), status)
}

func initRoutes() []router.Route {
	return []router.Route{
		router.NewGetRoute("/assetbundle/official/IOS/version", true, true, AssetBundleVersionGetRouteIOS),
		router.NewGetRoute("/assetbundle/official/Android/version", true, true, AssetBundleVersionGetRouteANDROID),
	}
}