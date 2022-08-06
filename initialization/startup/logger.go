package startup

import (
	"leafwearz/pkg/logging"
	"sync"

	"go.uber.org/zap"
)

type Logger struct {
	Required bool `json:"logging_required" yaml:"logging_required"`
	ServiceName string `json:"for_service" yaml:"for_service"`
	Directory string `json:"log_directory" yaml:"log_directory"`
	StorageMaxDay int `json:"storage_max_day" yaml:"storage_max_day"`
}

var LoggerGlobalSetting *Logger
var OperationLogger *zap.Logger
var HookLogger *zap.Logger

func InitLogger(wg *sync.WaitGroup) {
	defer wg.Done()
	operationCores := logging.GetCores("operation", LoggerGlobalSetting.Directory, LoggerGlobalSetting.StorageMaxDay)
	hookCores := logging.GetCores("hook", LoggerGlobalSetting.Directory, LoggerGlobalSetting.StorageMaxDay)
	caller := zap.AddCaller()
	panicOnlyMode := zap.Development()
	zapServiceName := zap.Fields(zap.String("ServiceName", LoggerGlobalSetting.ServiceName))
	OperationLogger = zap.New(operationCores, caller, panicOnlyMode, zapServiceName)
	HookLogger = zap.New(hookCores, caller, panicOnlyMode, zapServiceName)
	OperationLogger.Info("Initialization of Operation Zap Logger Finished...")
	OperationLogger.Info("Initialization of Hook Zap Logger Finished...")
	return
}