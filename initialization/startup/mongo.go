package startup

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDatabase struct {
	URI string `json:"mongo_uri" yaml:"mongo_uri"`
	DBName string `json:"mongo_db_name" yaml:"mongo_db_name"`
	AuthSource string `json:"mongo_auth_source" yaml:"mongo_auth_source"`
	User string `json:"mongo_user" yaml:"mongo_user"`
	Password string `json:"mongo_password" yaml:"mongo_password"`
	PoolLimit int `json:"mongo_pool_limit" yaml:"mongo_pool_limit"`
}

var MongoGlobalSetting *MongoDatabase
var MongoClient *mongo.Database

func InitMongo(wg *sync.WaitGroup) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	credential := options.Credential{
		AuthSource: MongoGlobalSetting.AuthSource,
		Username: MongoGlobalSetting.User,
		Password: MongoGlobalSetting.Password,
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoGlobalSetting.URI).SetAuth(credential))
	if err != nil {
		panic(err)
	}
	// check connection
	ctxPing, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctxPing, readpref.Primary()); err != nil {
        panic(err)
	}
	// bind client to global instance variables
	MongoClient = client.Database(MongoGlobalSetting.DBName)
}