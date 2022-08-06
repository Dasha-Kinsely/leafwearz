package startup

import (
	"leafwearz/globals"
	"leafwearz/initialization/migration"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySql(wg *sync.WaitGroup) {
	defer wg.Done()
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       globals.MySqlGlobalSetting.DSN,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  globals.MySqlGlobalSetting.DateTimePrecision,
		DontSupportRenameIndex:    globals.MySqlGlobalSetting.RenameIndex,
		DontSupportRenameColumn:   globals.MySqlGlobalSetting.RenameColumn,
		SkipInitializeWithVersion: globals.MySqlGlobalSetting.SkipInitializeWithVersion,
	}), &gorm.Config{})
	if err != nil {
		panic("Error initializing MYSQL_DB...")
	}
	if globals.ServerGlobalSetting.Migrate {
		if err := migration.BulkMigrate(); err != nil {
			panic("Failed to Migrate....")
		}
	}
	globals.MySqlClient = db
	return
}