package startup

import (
	"errors"
	"flag"
	"io/ioutil"
	"leafwearz/globals"
	"sync"

	"github.com/go-yaml/yaml"
)

func LoadConfigFiles() {
	// Get ConfigFiles' Paths
	confServerPath := flag.String("Srv", "initialization/configure/server.conf.yaml", "server config file path")
	confSqlPath := flag.String("SDB", "initialization/configure/mysql.conf.yaml", "sql db config file path")
	confNoSqlPath := flag.String("NSDB", "initialization/configure/mongo.conf.yaml", "nosql db config file path")
	confLoggerPath := flag.String("LOG", "initialization/configure/logger.conf.yaml", "logger config file path")
	confCorsPath := flag.String("Cors", "initialization/configure/cors.conf.yaml", "cors config file path")
	confRedisPath := flag.String("Redis", "initialization/configure/redis.conf.yaml", "redis conf file path")
	flag.Parse()
	// Load and Bind ConfigFiles into Global Objects
	var wg sync.WaitGroup
	wg.Add(6)
	go LoadConf(*confServerPath, "Server", &wg)
	go LoadConf(*confCorsPath, "Cors", &wg)
	go LoadConf(*confSqlPath, "MySql", &wg)
	go LoadConf(*confNoSqlPath, "Mongo", &wg)
	go LoadConf(*confLoggerPath, "Logger", &wg)
	go LoadConf(*confRedisPath, "Redis", &wg)
	defer wg.Wait()
}

func LoadConf(path string, configType string, wg *sync.WaitGroup) {
	defer wg.Done()
	data, err := ioutil.ReadFile(path)
	switch configType {
	case "Server" :
		err = yaml.Unmarshal(data, &globals.ServerGlobalSetting)
	case "MySql":
		err = yaml.Unmarshal(data, &globals.MySqlGlobalSetting)
	case "Mongo":
		err = yaml.Unmarshal(data, &globals.MongoGlobalSetting)
	case "Logger":
		err = yaml.Unmarshal(data, &globals.LoggerGlobalSetting)
	case "Cors":
		err = yaml.Unmarshal(data, &globals.CorsGlobalSetting)
	case "Redis":
		err = yaml.Unmarshal(data, &globals.RedisGlobalSetting)
	default:
		err = errors.New("Invalid Config Type! Failed to Bind Config-File into Objects")
	}
	if err != nil {
		panic("Error Occurred while Attempting to Read and Bind Config-Files into Objects... Fatal Error: Cannot Handle Further Requests !!!")
	}
	return
}
