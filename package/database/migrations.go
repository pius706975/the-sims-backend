package database

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" 
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/spf13/cobra"
)

var MigrationCMD = &cobra.Command{
	Use:   "migration",
	Short: "DB migration",
	RunE:  dbMigrate,
}

var migUp bool
var migDown bool

func init() {
	MigrationCMD.Flags().BoolVarP(&migUp, "up", "u", true, "run migration up")

	MigrationCMD.Flags().BoolVarP(&migDown, "down", "d", false, "run migration down")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	dsn := GetDatabaseURL()

	m, err := migrate.New(
		"file://package/database/migrations",
		dsn,
	)
	if err != nil {
		return err
	}

	if migUp {
		log.Println("Running migration up...")
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			return err
		}
		log.Println("Migration up done")
	}

	if migDown {
		log.Println("Running migration down...")
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			return err
		}
		log.Println("Migration down done")
	}

	return nil
}
