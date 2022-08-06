package startup

import (
	"leafwearz/globals"
	"leafwearz/pkg/logging"
	"sync"

	"go.uber.org/zap"
)

func InitLogger(wg *sync.WaitGroup) {
	defer wg.Done()
	operationCores := logging.GetCores("operation", globals.LoggerGlobalSetting.Directory, globals.LoggerGlobalSetting.StorageMaxDay)
	hookCores := logging.GetCores("hook", globals.LoggerGlobalSetting.Directory, globals.LoggerGlobalSetting.StorageMaxDay)
	caller := zap.AddCaller()
	panicOnlyMode := zap.Development()
	zapServiceName := zap.Fields(zap.String("ServiceName", globals.LoggerGlobalSetting.ServiceName))
	globals.OperationLogger = zap.New(operationCores, caller, panicOnlyMode, zapServiceName)
	globals.HookLogger = zap.New(hookCores, caller, panicOnlyMode, zapServiceName)
	globals.OperationLogger.Info("Initialization of Operation Zap Logger Finished...")
	globals.OperationLogger.Info("Initialization of Hook Zap Logger Finished...")
	return
}