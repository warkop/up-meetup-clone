package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "strconv"

	"github.com/warkop/up-meetup-clone/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	MigrateHandler = func(cmd *cobra.Command, args []string) {
		conf := config.LoadDatabaseConfig()

		goose.SetDialect(conf.Engine)

		inst, err := sql.Open(
			conf.Engine,
			fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=local",
				conf.User,
				conf.Password,
				conf.Host,
				conf.Port,
				conf.Schema,
			),
		)

		if err != nil {
			log.Fatal(err)
		}

		dir := viper.GetString("root_dir") + "/database/migration"

		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		appName := args[0]
		appArgs := args[1:]

		if err := goose.Run(appName, inst, dir, appArgs...); err != nil {
			log.Fatalf("(goose error): %v\n", err.Error())
		}
	}

	MigrateCommand = &cobra.Command{
		Use:   "migrate",
		Short: "Run database migrations.",
		Run:   MigrateHandler,
	}

	UsageCommand = `
Run database migrations

Usage:
	user migrate [command]

Available Commands:
	up                   Migrate the DB to the most recent version available.
	up-to VERSION        Migrate the DB to a specific version.
	down                 Rollback the version by 1.
	down-to VERSION      Rollback to a specific VERSION.
	redo                 Re-run the latest migration.
	status               Dump the migration status for the current DB.
	version              Print the current version of the database.
	create NAME [sql|go] Creates new migration file with next version.
	`
)

func InitializeMigrationCommand() {
	MigrateCommand.SetHelpFunc(func(*cobra.Command, []string) {
		log.Print(UsageCommand)
	})

	rootCmd.AddCommand(MigrateCommand)
}
