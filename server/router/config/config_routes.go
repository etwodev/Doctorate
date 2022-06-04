package config

import (
	"net/http"
	"github.com/Etwodev/Doctorate/server/helpers"
	"github.com/Etwodev/Doctorate/static"
)

func NetworkConfigGetRoute(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("sign") == "" {
		helpers.RespondWithError(w, http.StatusBadRequest, "No sign attribute")
	} else {
		helpers.RespondWithJSON(w, http.StatusOK, static.NetworkConfig, "application/octet-stream")
	}
}

func RemoteConfigGetRoute(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithCode(w, http.StatusOK, "200")
}