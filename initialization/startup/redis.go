package startup

import (
	"leafwearz/globals"
	"sync"
	"time"

	"github.com/go-redis/redis/v9"
)

func InitRedis(wg *sync.WaitGroup) {
	defer wg.Done()
	if globals.RedisGlobalSetting.CachingRequired {
		globals.RedisClient = redis.NewClient(&redis.Options{
			Addr: globals.RedisGlobalSetting.Addr,
			Password: globals.RedisGlobalSetting.Password,
			DB: globals.RedisGlobalSetting.DB,
			PoolSize: globals.RedisGlobalSetting.PoolSize,
			MinIdleConns: globals.RedisGlobalSetting.MinIdleConn,
			MaxRetries: 3,
			DialTimeout: 5*time.Second,
			ReadTimeout: 5*time.Second,
			WriteTimeout: 5*time.Second,
		})
	}
	return
}