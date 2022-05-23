package app

import (
	"encoding/json"
	"net/http"

	"github.com/Etwodev/Doctorate/server/helpers"
)

func AppCodesPostRoute(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("codestr")
	all := r.URL.Query().Get("all")
	var table Codes
	
	bin, err := helpers.OpenFile("./static/config/Codes.json")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	err = json.Unmarshal(bin, &table)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	if all == "1" || code == "" || all == "" {
		helpers.RespondWithRaw(w, http.StatusOK, bin, "application/json")
		return
	}

	t := []Code{}
	for _, c := range table.Data {
		if c.Number == code {
			t = append(t, c)
		}
	}

	table.Data = t
	helpers.RespondWithJSON(w, http.StatusOK, &table, "application/json")
}

func AppSettingsPostRoute(w http.ResponseWriter, r *http.Request) {
	bin, err := helpers.OpenFile("./static/config/Settings.json")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}
	helpers.RespondWithRaw(w, http.StatusOK, bin, "application/json")
}