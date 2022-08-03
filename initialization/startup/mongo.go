package startup

type MongoDatabase struct {
	URI string `json:"mongo_uri" yaml:"mongo_uri"`
	DBName string `json:"mongo_db_name" yaml:"mongo_db_name"`
	AuthSource string `json:"mongo_auth_source" yaml:"mongo_auth_source"`
	User string `json:"mongo_user" yaml:"mongo_user"`
	Password string `json:"mongo_password" yaml:"mongo_password"`
	PoolLimit int `json:"mongo_pool_limit" yaml:"mongo_pool_limit"`
}

var MongoGlobalSetting *MongoDatabase

