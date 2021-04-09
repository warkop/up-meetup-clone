package cmd

import (
	"context"
	"os"
	"os/signal"
	"time"

	_ "github.com/warkop/up-meetup-clone/config"
	_ "github.com/warkop/up-meetup-clone/logger"
	"github.com/warkop/up-meetup-clone/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "boot",
	Short: "Boot user http service.",
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()

		e.Pre(middleware.RemoveTrailingSlash())

		/*e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: logger.MiddlewareLog,
		}))*/

		// handle JWT
		e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey: []byte("fill-signing-key-here"),
		}))

		// handle CORS
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowHeaders: []string{
				echo.HeaderOrigin,
				echo.HeaderContentType,
				echo.HeaderAccept,
				echo.HeaderAuthorization,
			},
		}))

		routes.Endpoints(e)

		go func() {
			if err := e.Start(":" + viper.GetString("http.port")); err != nil {
				e.Logger.Info("Shutting down the server.")
			}
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)

		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	},
}

func InitializeBootCommand() {
	rootCmd.AddCommand(startCmd)
}
