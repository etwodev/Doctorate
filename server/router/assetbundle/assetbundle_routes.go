package assetbundle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Etwodev/Doctorate/server/helpers"
	stat "github.com/Etwodev/Doctorate/static"
	"github.com/Etwodev/Doctorate/types"
	"github.com/go-chi/chi/v5"
)

func AssetBundleVersionGetRoute(w http.ResponseWriter, r *http.Request) {
	dev := chi.URLParam(r, "device")
	var p types.Payload
	var v types.Version

	if (dev == "IOS" || dev == "Android") {
		bin, err := helpers.OpenFile("./static/hotupdate/" + dev + "/hot_update_list.json")
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
			return
		}

		err = json.Unmarshal(bin, &p)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
			return
		}

		v.Client = stat.ClientVersion
		v.Resource = p.VersionID
		
		helpers.RespondWithJSON(w, http.StatusOK, v, "application/octet-stream")
		return 
	} else {
		helpers.RespondWithError(w, http.StatusNotFound, "Page not found")
	}
}


func AssetBundleHotVersionGetRoute(w http.ResponseWriter, r *http.Request) {
	dev := chi.URLParam(r, "device")
	ver := chi.URLParam(r, "version")

	if (dev == "IOS" || dev == "Android") && ver != "" {
		path := filepath.Clean(strings.ReplaceAll(r.URL.Path, fmt.Sprintf("/assetbundle/official/%s/assets/%s/", dev, ver), ""))
		path = fmt.Sprintf("./static/hotupdate/%s", path)
		bin, err := helpers.OpenFile(path)
		if err != nil {
			helpers.RespondWithError(w, http.StatusNotFound, "Page not found")
			return
		}
		helpers.RespondWithRaw(w, http.StatusOK, bin, "application/octet-stream")
	} else {
		helpers.RespondWithError(w, http.StatusNotFound, "Page not found")
	}
}