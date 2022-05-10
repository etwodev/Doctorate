package announce

import (
	"net/http"

	"github.com/Etwodev/Doctorate/server/httputils"
	"github.com/Etwodev/Doctorate/server/router"
)

func AnnouncementPostRoute(w http.ResponseWriter, r *http.Request) {
	var ann []Announcement
	router.Connector.Find(&ann)
	// Cleanup this, find definitions for Extra and Focus
	cer := &Announcers{
		Announcements: ann,
		Extra: 		   &Xtrannounce{Enable: false, Name: "额外活动"},
		Focus:		   "480",
	}
	httputils.RespondWithJSON(w, http.StatusOK, cer)
}

func PreannouncementPostRoute(w http.ResponseWriter, r *http.Request) {
	var pre Preannouncement
	router.Connector.Find(&pre)
	httputils.RespondWithJSON(w, http.StatusOK, pre)
}

func VersionGetRoute(w http.ResponseWriter, r *http.Request) {
	var ver Version
	router.Connector.Find(&ver)
	httputils.RespondWithJSON(w, http.StatusOK, ver)
}
