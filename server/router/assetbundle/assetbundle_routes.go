package assetbundle

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Etwodev/Doctorate/server/helpers"
	"github.com/Etwodev/Doctorate/static"
	"github.com/go-chi/chi/v5"
)

func AssetBundleVersionGetRoute(w http.ResponseWriter, r *http.Request) {
	switch chi.URLParam(r, "device") {
	case "IOS":
		helpers.RespondWithJSON(w, http.StatusOK, static.IOS_VERSION, "application/octet-stream")
	case "Android":
		helpers.RespondWithJSON(w, http.StatusOK, static.ANDROID_VERSION, "application/octet-stream")
	default:
		helpers.RespondWithError(w, http.StatusNotFound, "Page not found")
	} 
}


func AssetBundleHotVersionGetRoute(w http.ResponseWriter, r *http.Request) {
	version := chi.URLParam(r, "version")
	device := chi.URLParam(r, "device")
	var path string

	if version == static.IOS_VERSION.Resource && device == "IOS" {
		path = fmt.Sprintf("/assetbundle/IOS/assets/%s/", version)
	}

	if version == static.ANDROID_VERSION.Resource && device == "Android" {
		path = fmt.Sprintf("/assetbundle/Android/assets/%s/", version)
	}

	path = fmt.Sprintf("./static/hotupdate/%s", filepath.Clean(strings.ReplaceAll(r.URL.Path, path, "")))
	if path == "" {
		helpers.RespondWithError(w, http.StatusNotFound, "Page not found")
	}

	bin, err := helpers.OpenFile(path)
	if err != nil {
		helpers.RespondWithError(w, http.StatusNotFound, "File not found")
	} 

	helpers.RespondWithRaw(w, http.StatusOK, bin, "application/octet-stream")
}