package config

import (
	"net/http"
	"github.com/Etwodev/Doctorate/server/helpers"
)

func NetworkConfigGetRoute(w http.ResponseWriter, r *http.Request) {
	sign := r.URL.Query().Get("sign")
	if sign == "" {
		helpers.RespondWithError(w, http.StatusBadRequest, "No sign attribute")
		return
	}

	content, err := helpers.Serialization("./static/config/NetworkConfig.json")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	resign, err := helpers.SignatureWithMD5(content, "./static/keys/private.key")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	response := NetworkConfig {
		Sign: resign,
		Content: content,
	}

	helpers.RespondWithJSON(w, http.StatusOK, response, "application/octet-stream")
}

func RemoteConfigGetRoute(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithCode(w, http.StatusOK, "200")
}