package commands

import (
	"github.com/spf13/cobra"

	"github.com/juanfer2/golang-clean/src/shared/infrastructure/servers"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server about this service",
	Long:  `server about this service`,
	Run:   runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	servers.StartServerApp()
}
