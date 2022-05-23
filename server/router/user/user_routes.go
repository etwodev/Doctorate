package user

import (
	"net/http"

	"github.com/Etwodev/Doctorate/server/helpers"
)


func AgreementPostRoute(w http.ResponseWriter, r *http.Request) {
	bin, err := helpers.OpenFile("./static/config/Agreements.json")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}
	helpers.RespondWithRaw(w, http.StatusOK, bin, "application/json")
}