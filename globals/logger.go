package globals

type Logger struct {
	Required bool `json:"logging_required" yaml:"logging_required"`
	ServiceName string `json:"for_service" yaml:"for_service"`
	Directory string `json:"log_directory" yaml:"log_directory"`
	StorageMaxDay int `json:"storage_max_day" yaml:"storage_max_day"`
}