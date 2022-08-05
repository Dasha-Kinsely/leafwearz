package startup

import (
	"leafwearz/routers"
)

func DefaultRun() {
	// if anything goes wrong at this step, terminal by panicing
	LoadConfigFiles() 
	/* To avoid data race, you must load Config-Files before proceeding with the actual initialization of:
		1) databases, 2) loggers, 3) caching, and 4) server-with-middlewares */
	InitAll()
	AddCommonMiddleware()
	// Initialize routers with controller interfaces and structs after everything else is ready, then run the server
	DefaultServer = routers.InitRouterGroup(DefaultServer)
	DefaultServer.Run()
}
