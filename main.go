package main

import (
	"up-meetup-clone/cmd"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// login is contract for login
type login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func setupEnv() {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")

}

func main() {
	gin.ForceConsoleColor()
	r := gin.Default()

	setupEnv()
	err := viper.ReadInConfig()

	if err != nil {
		r.Use(gin.Logger())
	}

	port := viper.GetString("SERVER_PORT")

	if port == "" {
		port = "8000"
	}

	api := r.Group("/api")

	api.POST("/auth", cmd.Auth)

	title := api.Group("/title")
	// title.Use(AuthRequired())
	title.GET("/", cmd.ListTitle)
	title.GET("/:id", cmd.DetailTitle)

	r.Run(":" + port)
}
