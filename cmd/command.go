package cmd

import (
	serve "github.com/pius706975/the-sims-backend/api"
	"github.com/pius706975/the-sims-backend/cmd/createsuperuser"
	"github.com/pius706975/the-sims-backend/package/database"

	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "backend",
	Long: "Go backend service",
}

func init() {
	initCommand.AddCommand(serve.ServeCMD)
	initCommand.AddCommand(database.MigrationCMD)
	initCommand.AddCommand(createsuperuser.CreateSuperUserCMD)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}