package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	_ "github.com/warkop/up-meetup-clone/config"
	"github.com/warkop/up-meetup-clone/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "boot",
	Short: "Boot user http service.",
	Run: func(cmd *cobra.Command, args []string) {
		app := fiber.New(fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(500).SendString(err.Error())
			},
		})

		app.Use(fiberLogger.New(fiberLogger.Config{
			Next:         nil,
			Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
			TimeFormat:   "15:04:05",
			TimeZone:     "Local",
			TimeInterval: 500 * time.Millisecond,
			Output:       os.Stderr,
		}))

		// handle JWT
		// app.Use(jwtware.New(jwtware.Config{
		// 	SigningKey: []byte("my Secret key!"),
		// }))

		// handle CORS
		app.Use(cors.New(cors.Config{
			AllowHeaders: fmt.Sprintf(`%s, %s, %s, %s`,
				fiber.HeaderOrigin,
				fiber.HeaderContentType,
				fiber.HeaderAccept,
				fiber.HeaderAuthorization,
			),
		}))

		routes.Endpoints(app)

		go func() {
			if err := app.Listen(":" + viper.GetString("http.port")); err != nil {
				app.Server().Logger.Printf("Shutting down the server.")
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)

		<-quit

		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := app.Shutdown(); err != nil {
			app.Server().Logger.Printf(`%s`, err)
		}
	},
}

func InitializeBootCommand() {
	rootCmd.AddCommand(startCmd)
}
