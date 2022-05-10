package announce

import (
	"github.com/Etwodev/Doctorate/server/router"
)

func NewRouter(status bool) router.Router {
	router.Connector.AutoMigrate(&Announcement{}, &Preannouncement{}, &Version{})
	return router.NewRouter(initRoutes(), status)
}

func initRoutes() []router.Route {
	return []router.Route{
		router.NewGetRoute("/announce/announcements", true, true, AnnouncementPostRoute),
		router.NewGetRoute("/announce/preannouncements", true, true, PreannouncementPostRoute),
		router.NewGetRoute("/announce/version", true, true, VersionGetRoute),
	}
}