package startup

import (
	"context"
	"leafwearz/globals"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongo(wg *sync.WaitGroup) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	credential := options.Credential{
		AuthSource: globals.MongoGlobalSetting.AuthSource,
		Username: globals.MongoGlobalSetting.User,
		Password: globals.MongoGlobalSetting.Password,
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(globals.MongoGlobalSetting.URI).SetAuth(credential))
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
	globals.MongoClient = client.Database(globals.MongoGlobalSetting.DBName)
	return
}