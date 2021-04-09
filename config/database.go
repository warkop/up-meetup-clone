package config

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Orm *gorm.DB
)

type Database struct {
	Engine            string
	Host              string
	User              string
	Password          string
	Schema            string
	Port              int
	ReconnectRetry    int
	ReconnectInterval int64
	DebugMode         bool
}

func LoadDatabaseConfig() Database {
	conf := Database{
		Engine:            viper.GetString("database.engine"),
		Host:              viper.GetString("database.host"),
		User:              viper.GetString("database.username"),
		Password:          viper.GetString("database.password"),
		Schema:            viper.GetString("database.schema"),
		Port:              viper.GetInt("database.port"),
		ReconnectRetry:    viper.GetInt("database.reconnect_retry"),
		ReconnectInterval: viper.GetInt64("database.reconnect_interval"),
		DebugMode:         viper.GetBool("database.debug_mode"),
	}

	return conf
}

func DBConnect() *gorm.DB {
	conf := LoadDatabaseConfig()
	inst, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                        // default size for string fields
		DisableDatetimePrecision:  true,                                                                       // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                       // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                       // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                      // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// inst.LogMode(conf.DebugMode)

	if conf.DebugMode {
		return inst.Debug()
	}

	return inst
}

func OpenDB() {
	Orm = DBConnect()
}
