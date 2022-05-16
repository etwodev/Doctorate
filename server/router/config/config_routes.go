package config

import (
	"net/http"
	"github.com/Etwodev/Doctorate/server/helpers"
)

func NetworkConfigGetRoute(w http.ResponseWriter, r *http.Request) {
	sign := r.URL.Query().Get("sign")
	if sign == "" {
		helpers.RespondWithError(w, 400, "No sign attribute")
		return
	}

	resign := helpers.MD5SignWithPrivateKey(helpers.PrivKey, helpers.HostConfigSTR)

	res := NetworkConfig {
		Sign: resign,
		Content: helpers.HostConfigRAW,
	}

	helpers.RespondWithRawJSON(w, http.StatusOK, res)
}

func RemoteConfigGetRoute(w http.ResponseWriter, r *http.Request) {
	// FIXME: Establish connection with client and mimic servers, currently just freezes
	var res interface{}
	helpers.RespondWithJSON(w, http.StatusOK, res)
}