package startup

import (
	"leafwearz/initialization/migration"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlDatabase struct {
	DSN string `json:"mysql_dsn" yaml:"mysql_dsn"`
	DateTimePrecision bool `json:"mysql_disable_datetime_precision" yaml:"mysql_disable_datetime_precision"`
	RenameColumn bool `json:"mysql_dont_support_rename_column" yaml:"mysql_dont_support_rename_column"`
	RenameIndex bool `json:"mysql_dont_support_rename_index" yaml:"mysql_dont_support_rename_index"`
	MaxOpenConn int `json:"mysql_max_open_conn" yaml:"mysql_max_open_conn"`
	MaxIdleConn int `json:"mysql_max_idle_conn" yaml:"mysql_max_idle_conn"`
	SkipInitializeWithVersion bool `json:"mysql_skip_initialize_with_version" yaml:"mysql_skip_initialize_with_version"`
}

var MySqlGlobalSetting *MySqlDatabase
var MySqlClient *gorm.DB

func InitMySql(wg *sync.WaitGroup) {
	defer wg.Done()
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       MySqlGlobalSetting.DSN,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  MySqlGlobalSetting.DateTimePrecision,
		DontSupportRenameIndex:    MySqlGlobalSetting.RenameIndex,
		DontSupportRenameColumn:   MySqlGlobalSetting.RenameColumn,
		SkipInitializeWithVersion: MySqlGlobalSetting.SkipInitializeWithVersion,
	}), &gorm.Config{})
	if err != nil {
		panic("Error initializing MYSQL_DB...")
	}
	if ServerGlobalSetting.Migrate {
		if err := migration.BulkMigrate(); err != nil {
			panic("Failed to Migrate....")
		}
	}
	MySqlClient = db
}