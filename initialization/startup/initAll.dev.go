package startup

import (
	"sync"
)

func InitAll() {
	var wg sync.WaitGroup
	wg.Add(5)
	go InitMySql(&wg)
	go InitMongo(&wg)
	go InitRedis(&wg)
	go InitServer(&wg)
	go InitResources(&wg)
	// go InitKafka(&wg)
	defer wg.Wait()
	// InitLogger()
}