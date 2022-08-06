package startup

import (
	"sync"
)

func InitAll() {
	var wg sync.WaitGroup
	wg.Add(4)
	go InitMySql(&wg)
	go InitMongo(&wg)
	go InitRedis(&wg)
	go InitServer(&wg)
	// go InitKafka(&wg)
	defer wg.Wait()
	// InitLogger()
}