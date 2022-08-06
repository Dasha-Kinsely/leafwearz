package globals

type MySqlDatabase struct {
	DSN string `json:"mysql_dsn" yaml:"mysql_dsn"`
	DateTimePrecision bool `json:"mysql_disable_datetime_precision" yaml:"mysql_disable_datetime_precision"`
	RenameColumn bool `json:"mysql_dont_support_rename_column" yaml:"mysql_dont_support_rename_column"`
	RenameIndex bool `json:"mysql_dont_support_rename_index" yaml:"mysql_dont_support_rename_index"`
	MaxOpenConn int `json:"mysql_max_open_conn" yaml:"mysql_max_open_conn"`
	MaxIdleConn int `json:"mysql_max_idle_conn" yaml:"mysql_max_idle_conn"`
	SkipInitializeWithVersion bool `json:"mysql_skip_initialize_with_version" yaml:"mysql_skip_initialize_with_version"`
}