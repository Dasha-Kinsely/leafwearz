package globals

type Redis struct {
	CachingRequired bool `json:"caching_required" yaml:"caching_required"`
	Addr string `json:"redis_addr" yaml:"redis_addr"`
	Password string `json:"redis_password" yaml:"redis_password"`
	DB int `json:"redis_db" yaml:"redis_db"`
	PoolSize int `json:"redis_pool_size" yaml:"redis_pool_size"`
	MinIdleConn int `json:"redis_min_idle_conn" yaml:"redis_min_idle_conn"`
}