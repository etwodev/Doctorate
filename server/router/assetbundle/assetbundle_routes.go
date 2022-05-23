package assetbundle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Etwodev/Doctorate/server/helpers"
	"github.com/go-chi/chi/v5"
)

func AssetBundleVersionGetRoute(w http.ResponseWriter, r *http.Request) {
	dev := chi.URLParam(r, "device")
	var ver Versions

	bin, err := helpers.OpenFile("./static/config/Versions.json")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	err = json.Unmarshal(bin, &ver)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	switch dev {
	case "IOS":
		helpers.RespondWithJSON(w, http.StatusOK, ver.IOS, "application/octet-stream")
	case "Android":
		helpers.RespondWithJSON(w, http.StatusOK, ver.Android, "application/octet-stream")
	default:
		helpers.RespondWithError(w, http.StatusNotFound, "Page not found")
	}
}

func AssetBundleHotVersionGetRoute(w http.ResponseWriter, r *http.Request) {
	dev := chi.URLParam(r, "device")
	if (dev == "IOS" || dev == "Android") {
		path := filepath.Clean(strings.ReplaceAll(r.URL.Path, fmt.Sprintf("/assetbundle/official/%s/assets", dev), ""))
		path = fmt.Sprintf("./static/hotupdate/%s%s", dev, path)
		bin, err := helpers.OpenFile(path)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
			return
		}
		helpers.RespondWithRaw(w, http.StatusOK, bin, "application/octet-stream")
	} else {
		helpers.RespondWithError(w, http.StatusNotFound, "Page not found")
	}
}