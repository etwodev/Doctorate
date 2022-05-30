package announce

import (
	"net/http"

	"github.com/Etwodev/Doctorate/server/helpers"
)


func PreannounceGetRoute(w http.ResponseWriter, r *http.Request) {
	bin, err := helpers.OpenFile("./static/config/Preannouncement.json")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}
	helpers.RespondWithRaw(w, http.StatusOK, bin, "application/json")
}