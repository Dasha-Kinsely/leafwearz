package startup

import (
	"errors"
	"flag"
	"io/ioutil"
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
	flag.Parse()
	// Load and Bind ConfigFiles into Global Objects
	var wg sync.WaitGroup
	wg.Add(5)
	go LoadConf(*confServerPath, "Server", &wg)
	go LoadConf(*confCorsPath, "Cors", &wg)
	go LoadConf(*confSqlPath, "MySql", &wg)
	go LoadConf(*confNoSqlPath, "Mongo", &wg)
	go LoadConf(*confLoggerPath, "Logger", &wg)
	defer wg.Wait()
}

func LoadConf(configType string, path string, wg *sync.WaitGroup) {
	defer wg.Done()
	data, err := ioutil.ReadFile(path)
	switch configType {
	case "Server" :
		err = yaml.Unmarshal(data, &ServerGlobalSetting)
	case "MySql":
		err = yaml.Unmarshal(data, &MySqlGlobalSetting)
	case "Mongo":
		err = yaml.Unmarshal(data, &MongoGlobalSetting)
	case "Logger":
		err = yaml.Unmarshal(data, &LoggerGlobalSetting)
	case "Cors":
		err = yaml.Unmarshal(data, &CorsGlobalSetting)
	default:
		err = errors.New("Invalid Config Type! Failed to Bind Config-File into Objects")
	}
	if err != nil {
		panic("Error Occurred while Attempting to Read and Bind Config-Files into Objects... Fatal Error: Cannot Handle Further Requests !!!")
	}
}
