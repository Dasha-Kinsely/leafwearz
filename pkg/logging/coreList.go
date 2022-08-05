package logging

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetCores(option string, directory string, storageLength int) zapcore.Core {
	switch option {
	case "operation":
		core := zapcore.NewTee(
			getOperationCoreList(directory, storageLength)...,
		)
		return core
	case "hook":
		core := zapcore.NewTee(
			getHookCoreList(directory, storageLength)...
		)
		return core
	default:
		panic("Cannot Get this Type of Core")
	}
}

func addSync(w io.Writer) zapcore.WriteSyncer {
	return zapcore.AddSync(w)
}

func getOperationCoreList(directory string, storageLength int) []zapcore.Core {
	// Info Logs
	infoLoggerLevelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.InfoLevel
	})
	rotateInfoLogs := getRotateLogs(directory + "info.log", storageLength)
	// Warn Logs
	warnLoggerLevelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})
	rotateWarnLogs := getRotateLogs(directory + "warn.log", storageLength)
	// Kafka and Mongo Custom Logs
	kafkaWriter := &KafkaLogWriter{}
	mongoWriter := &MongoLogWriter{}
	// Append all into their respective WriterSyncLists
	var infoWriterSyncList, warnWriterSyncList []zapcore.WriteSyncer
	infoWriterSyncList = append(infoWriterSyncList, addSync(os.Stdout), addSync(rotateInfoLogs), addSync(kafkaWriter), addSync(mongoWriter)) 
	warnWriterSyncList = append(warnWriterSyncList, addSync(os.Stdout), addSync(rotateWarnLogs), addSync(kafkaWriter), addSync(mongoWriter))
	// return coreList
	var coreList []zapcore.Core
	coreList = append(coreList,
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(infoWriterSyncList...),
			infoLoggerLevelEnabler,
		),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(warnWriterSyncList...),
			warnLoggerLevelEnabler,
		))
	return coreList
}

func getHookCoreList(directory string, storageLength int) []zapcore.Core {
	debugLoggerLevelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})
	rotateHookLogs := getRotateLogs(directory + "hook.log", storageLength)
	var hookWriterSyncList []zapcore.WriteSyncer
	hookWriterSyncList = append(hookWriterSyncList, addSync(os.Stdout), addSync(rotateHookLogs))
	var coreList []zapcore.Core
	coreList = append(coreList,
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(hookWriterSyncList...),
			debugLoggerLevelEnabler,
		))
	return coreList
}