package customcmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

var migrationName string

var CreateMigrationCMD = &cobra.Command{
	Use:   "create-migration",
	Short: "Generate SQL migration files with version sequence",
	RunE:  createMigration,
}

func init() {
	CreateMigrationCMD.Flags().StringVar(&migrationName, "name", "", "Migration name")
	CreateMigrationCMD.MarkFlagRequired("name")
}

func createMigration(cmd *cobra.Command, args []string) error {
	migrationsDir := "package/database/migrations"

	// check if the directory exists
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		os.MkdirAll(migrationsDir, os.ModePerm)
	}

	// read the migration files
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	// find the last version number of the migration files
	version := 0
	r := regexp.MustCompile(`^(\d+)_`)
	for _, f := range files {
		match := r.FindStringSubmatch(f.Name())
		if len(match) > 1 {
			v, _ := strconv.Atoi(match[1])
			if v > version {
				version = v
			}
		}
	}

	// increment version number for the new migration file
	version++
	versionStr := fmt.Sprintf("%04d", version)

	// generate file name
	upFile := filepath.Join(migrationsDir, fmt.Sprintf("%s_%s.up.sql", versionStr, migrationName))
	downFile := filepath.Join(migrationsDir, fmt.Sprintf("%s_%s.down.sql", versionStr, migrationName))

	upContent := "-- +++ UP migration +++\n\n"
	downContent := "-- +++ DOWN migration +++\n\n"

	if err := os.WriteFile(upFile, []byte(upContent), 0644); err != nil {
		return err
	}
	if err := os.WriteFile(downFile, []byte(downContent), 0644); err != nil {
		return err
	}

	fmt.Println("Migration files created:")
	fmt.Println(upFile)
	fmt.Println(downFile)
	return nil
}
