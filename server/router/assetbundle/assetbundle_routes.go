package assetbundle

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Etwodev/Doctorate/server/helpers"
	"github.com/go-chi/chi/v5"
)


// ark-us-static-online.yo-star.com/assetbundle/official/IOS/assets/../../Android/assets/22-04-28-22-22-02-df57bb/hot_update_list.json

func AssetBundleVersionGetRoute(w http.ResponseWriter, r *http.Request) {
	dev := chi.URLParam(r, "device")
	switch dev {
	case "IOS":
		helpers.RespondWithFileJSON(w, http.StatusOK, "./server/router/assetbundle/assets/AssetBundleVersionIOS.json")
	case "Android":
		helpers.RespondWithFileJSON(w, http.StatusOK, "./server/router/assetbundle/assets/AssetBundleVersionANDROID.json")
	default:
		helpers.RespondWithError(w, 404, "Page not found")
	}
}

func AssetBundleHotVersionGetRoute(w http.ResponseWriter, r *http.Request) {
	dev := chi.URLParam(r, "device")
	if (dev == "IOS" || dev == "Android") {
		path := filepath.Clean(strings.ReplaceAll(r.URL.Path, fmt.Sprintf("/assetbundle/official/%s/assets", dev), ""))
		path = fmt.Sprintf("./public/%s%s", dev, path)
		helpers.RespondWithOctet(w, http.StatusOK, path)
	} else {
		helpers.RespondWithError(w, 404, "Page not found")
	}
}