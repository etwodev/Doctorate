# [WIP] Doctorate

## Current Features

- [x] Announcements
- [x] Mail
- [ ] Logging in
- [ ] Missions
- [ ] Registering
- [ ] Updates
- [ ] Gacha
- [ ] Events
- [ ] Items & Rewards
- [ ] Operations / Gameplay

## Startup

### Getting Started

When starting up the server, the client only needs to connect for a couple seconds, where then it is safe to disconnect after the update has been installed. This is because the game uses a system where it only stores the remote config file. If you are running on a public Doctorate server, you will get all the updates a normal user would get, this is because Doctorate actually requests the updates from Official Servers on behalf of you and passes back adjusted data!
 
```mermaid
sequenceDiagram
Client ->> Proxy: GET [HOST] /remote_config
Proxy ->> Server: GET [FAKEHOST] /remote_config
Server--> Client: Response [FAKEHOST] {"gs":"[FAKEHOST]", "as":"[FAKEHOST]"...
Note right of Server: Client now thinks that<br> the [FAKEHOST] is the<br> real host, and sends all<br> requests to that server.
Client ->> Server: GET [FAKEHOST] /announce/version
Server--> Client: Reponse [FAKEHOST] {"resVersion":"22-04-21-10-53-07-1be683"...
Note right of Server: Server sends a version that<br> is greater than the clients<br> current version, causing<br> an update to be requested,<br> which in return, changes<br> the default host.
Client ->> Server: GET [FAKEHOST] /update
Server--> Client: Response [FAKEHOST] *Update Patch*
Client ->> Server: POST [FAKEHOST] /user/login
```

## Creating a new Router

Doctorate uses a module system to register new API routes. Each router "package" is enclosed inside the router. For example, say the game has a new path called "shop", this would be the folder name and then any routes under it are registered accordingly.

```mermaid
graph  TD
A[Router]
A --> |/api/shop| B(package shop)
A --> |/api/announce| C(package announce)
A --> |/api/quest| D(package quest)
```
We can create a new package by simply doing something like the following:
```go
// '/router/test/test.go'
package test

// Import the router package
import (
	"github.com/Etwodev/Doctorate/server/router"
) 

func  NewRouter(status bool) router.Router {
	// Automatically creates missing tables and more for us
	router.Connector.AutoMigrate(&Test{})
	return router.NewRouter(initRoutes(), status)
}

func  initRoutes() []router.Route {
	return []router.Route{
	// path, isExperimental, isActivated, method
		router.NewGetRoute("/test/test", true, true, TestPostRoute),
	}
}
```
TestPostRoute may look something like this:
```go
// '/router/test/test_routes.go
package test

import (
	"net/http"
	"github.com/Etwodev/Doctorate/server/httputils"
	"github.com/Etwodev/Doctorate/server/router"
)

func  TestPostRoute(w http.ResponseWriter, r *http.Request) {
	var t Test
	// Gets all data relating to the object Test 
	router.Connector.Find(&t)
	httputils.RespondWithJSON(w, http.StatusOK, pre)
}
```

We need to make sure we update our main.go file, though! 
This is to make sure the router gets initialised and we don't get an empty router.
```go
// main.go
...
routers := []router.Router{
	// We call NewRouter(), where "true" is whether the route is should be activated
	test.NewRouter(true),
}

// Initilises all the routers
s.InitRouter(routers...)
// In this case, "true" is whether or not to run the server in experimental mode
var  r = s.InitRouters(true)
...
```

