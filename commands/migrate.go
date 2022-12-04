package commands

import (
	"github.com/juanfer2/golang-clean/commands/migrations"
	"github.com/spf13/cobra"
	// "github.com/juanfer2/ayenda_service/src/config"
	// "github.com/juanfer2/ayenda_service/src/models"
)

var migrateCmd = &cobra.Command{
	Use:     "migration",
	Aliases: []string{"m"},
	Short:   "Migrate database actions",
}

func init() {
	migrateCmd.AddCommand(migrations.CreateMigrationCmd, migrations.RunMigrationCmd,
		migrations.DownMigrationCmd)
}
