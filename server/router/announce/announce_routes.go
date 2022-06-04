package announce

import (
	"net/http"

	"github.com/Etwodev/Doctorate/server/helpers"
	"github.com/Etwodev/Doctorate/static"
)


func PreannounceGetRoute(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithJSON(w, http.StatusOK, static.Preannouncement, "application/json")
}
