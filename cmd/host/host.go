package host

import (
	"fmt"
	"scango/cmd"

	"github.com/spf13/cobra"
)

var HostCmd = &cobra.Command{
	Use:   "host",
	Short: "Find active hosts on the network",
	Long:  `Find active hosts on the network. You can search for a specific IP address or an IP range.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Check the possible subcommands with the \"scango host --help\".")
	},
}

func init() {
	cmd.RootCmd.AddCommand(HostCmd)
}
