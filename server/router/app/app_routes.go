package app

import (
	"net/http"
	"github.com/Etwodev/Doctorate/server/helpers"
	"github.com/Etwodev/Doctorate/static"
	"github.com/Etwodev/Doctorate/types"
)

func AppCodesPostRoute(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("codestr")
	all := r.FormValue("all")
	var response types.Codes

	if all != "1" {
		for _, c := range static.Codes {
			if code == c.Number {
				response.Data = append(response.Data, c)
				break
			}
		}
	} else {
		response.Data = static.Codes
	}

	response.Result = 0
	helpers.RespondWithJSON(w, http.StatusOK, response, "application/json")
}

func AppSettingsPostRoute(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithJSON(w, http.StatusOK, static.Settings, "application/json")
}