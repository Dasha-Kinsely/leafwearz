package startup

import (
	"sync"
	"time"

	"github.com/go-redis/redis/v9"
)

type Redis struct {
	CachingRequired bool `json:"caching_required" yaml:"caching_required"`
	Addr string `json:"redis_addr" yaml:"redis_addr"`
	Password string `json:"redis_password" yaml:"redis_password"`
	DB int `json:"redis_db" yaml:"redis_db"`
	PoolSize int `json:"redis_pool_size" yaml:"redis_pool_size"`
	MinIdleConn int `json:"redis_min_idle_conn" yaml:"redis_min_idle_conn"`
}

var RedisGlobalSetting *Redis
var RedisClient *redis.Client

func InitRedis(wg *sync.WaitGroup) {
	defer wg.Done()
	if RedisGlobalSetting.CachingRequired {
		RedisClient = redis.NewClient(&redis.Options{
			Addr: RedisGlobalSetting.Addr,
			Password: RedisGlobalSetting.Password,
			DB: RedisGlobalSetting.DB,
			PoolSize: RedisGlobalSetting.PoolSize,
			MinIdleConns: RedisGlobalSetting.MinIdleConn,
			MaxRetries: 3,
			DialTimeout: 5*time.Second,
			ReadTimeout: 5*time.Second,
			WriteTimeout: 5*time.Second,
		})
	}
	return
}