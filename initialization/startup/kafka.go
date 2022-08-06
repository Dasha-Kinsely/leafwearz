package startup

import (
	"sync"
)

func InitKafka(wg *sync.WaitGroup) {
	defer wg.Done()
}