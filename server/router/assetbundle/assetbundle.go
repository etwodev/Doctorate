package assetbundle

import (
	"github.com/Etwodev/Doctorate/server/router"
)

func NewRouter(status bool) router.Router {
	return router.NewRouter(initRoutes(), status)
}

func initRoutes() []router.Route {
	return []router.Route{
		router.NewGetRoute("/assetbundle/{device}/version", true, true, AssetBundleVersionGetRoute),
		router.NewGetRoute("/assetbundle/{device}/assets/{version}/", true, true, AssetBundleHotVersionGetRoute),
	}
}