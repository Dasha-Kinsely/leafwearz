package startup

func DefaultRun() {
	InitRouterGroup()
	DefaultServer.Run()
}
