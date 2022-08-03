package startup

type Logger struct {
	Required bool `json:"logging_required" yaml:"logging_required"`
}

var LoggerGlobalSetting *Logger