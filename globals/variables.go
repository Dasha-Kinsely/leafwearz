package globals

import (
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	MySqlGlobalSetting *MySqlDatabase
	MySqlClient *gorm.DB
	MongoGlobalSetting *MongoDatabase
	MongoClient *mongo.Database
	RedisGlobalSetting *Redis
	RedisClient *redis.Client
	KafkaProducer sarama.SyncProducer
	ServerGlobalSetting *Server
	DefaultServer *gin.Engine
	CorsGlobalSetting *Cors
	LoggerGlobalSetting *Logger
	OperationLogger *zap.Logger
	HookLogger *zap.Logger
)