package port

import (
	"fmt"
	"scango/cmd"

	"github.com/spf13/cobra"
)

var PortCmd = &cobra.Command{
	Use:   "port",
	Short: "find open ports on a host",
	Long:  `find open ports on a host. You can search for a specific port or search for multiple ports.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Check the possible subcommands with the \"scango port --help\".")
	},
}

func init() {
	cmd.RootCmd.AddCommand(PortCmd)
}
