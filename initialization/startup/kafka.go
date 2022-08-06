package startup

import (
	"sync"

	"github.com/Shopify/sarama"
)

var KafkaProducer sarama.SyncProducer

func InitKafka(wg *sync.WaitGroup) {
	defer wg.Done()
}