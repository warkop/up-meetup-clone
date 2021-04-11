package config

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
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
	Port              string
	ReconnectRetry    int
	ReconnectInterval int64
	DebugMode         bool
	URL               string
}

func LoadDatabaseConfig() Database {
	conf := Database{
		Engine:            viper.GetString("database.engine"),
		Host:              viper.GetString("database.host"),
		User:              viper.GetString("database.username"),
		Password:          viper.GetString("database.password"),
		Schema:            viper.GetString("database.schema"),
		Port:              viper.GetString("database.port"),
		ReconnectRetry:    viper.GetInt("database.reconnect_retry"),
		ReconnectInterval: viper.GetInt64("database.reconnect_interval"),
		DebugMode:         viper.GetBool("database.debug_mode"),
		URL:               viper.GetString("database.url"),
	}

	return conf
}

func DBConnect() *gorm.DB {
	conf := LoadDatabaseConfig()
	inst, err := gorm.Open(postgres.Open(conf.URL), &gorm.Config{})

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
