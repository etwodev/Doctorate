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

	content := helpers.Serialization("./server/router/config/assets/NetworkConfigGetRoute.json")
	resign := helpers.MD5SignWithPrivateKey(content, "./keys/private.key")

	res := NetworkConfig {
		Sign: resign,
		Content: content,
	}

	helpers.RespondWithRawJSON(w, http.StatusOK, res)
}