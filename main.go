package main

import (
	"leafwearz/initialization/startup"
)

func main() {
	// If anything goes wrong at this step, terminal by panicing
	startup.LoadConfigFiles() 
	/* To avoid data race, you must load Config-Files before proceeding with the actual initialization of:
		1) databases, 2) server+middlewares, 3) kafka, 4) logger */
	startup.InitAll()
	// After everything else is ready, initialize routers with controller interfaces and structs, then run the server.
	startup.DefaultRun()
}