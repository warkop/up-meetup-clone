package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	AppPath string
)

func init() {
	LoadEnvVars()
	OpenDB()
}

func LoadEnvVars() {
	viper.SetEnvPrefix("majoo")
	viper.BindEnv("env")

	vpath := ""

	os.Setenv("MAJOO_ENV", "development")

	if viper.Get("env") == "development" {
		vpath = "dev"
	} else if viper.Get("env") == "production" {
		vpath = "prod"
	}

	cdir, _ := os.Getwd()

	viper.AddConfigPath(fmt.Sprintf("%s/config/%s", cdir, vpath))
	viper.SetConfigName("global.json")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err.Error()))
	}
}
