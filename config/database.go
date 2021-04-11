package config

import (
	"fmt"

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
	}

	return conf
}

func DBConnect() *gorm.DB {
	conf := LoadDatabaseConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", conf.Host, conf.User, conf.Password, conf.Schema, conf.Port)
	fmt.Println("dsn:", dsn)
	inst, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
