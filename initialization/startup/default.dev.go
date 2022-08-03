package startup

import (
	"leafwearz/routers"
)

func DefaultRun() {
	LoadConfigFiles() // if anything goes wrong at this step, terminal by panicing
	

	r := routers.InitRouterGroup()
	r.Run()
}
