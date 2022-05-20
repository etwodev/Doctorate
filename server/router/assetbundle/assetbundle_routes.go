package assetbundle

import (
	"net/http"

	"github.com/Etwodev/Doctorate/server/helpers"
)


func AssetBundleVersionGetRouteIOS(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithFileJSON(w, http.StatusOK, "./server/router/assetbundle/assets/AssetBundleVersionGetRouteIOS.json")
}
func AssetBundleVersionGetRouteANDROID(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithFileJSON(w, http.StatusOK, "./server/router/assetbundle/assets/AssetBundleVersionGetRouteANDROID.json")
}