package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect is for creating connection
func Connect() *gorm.DB {
	dsn := "host=" + viper.GetString("DB_HOST") + " port=" + viper.GetString("DB_PORT") + " user=" + viper.GetString("DB_USER") + " dbname=" + viper.GetString("DB_NAME") + " password=" + viper.GetString("DB_PASS") + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
