package user

import (
	"net/http"
	"github.com/Etwodev/Doctorate/server/helpers"
	"github.com/Etwodev/Doctorate/static"
)


func AgreementPostRoute(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithJSON(w, http.StatusOK, static.Agreements, "application/json")
}
