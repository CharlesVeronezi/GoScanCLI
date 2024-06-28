package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd representa o comando base quando chamado sem nenhum subcomando
var RootCmd = &cobra.Command{
	Use:   "scango",
	Short: "ScanGo is an open-source CLI for scanning IP addresses and ports on networks.",
	Long: `ScanGo is an open-source CLI for scanning IP addresses and ports on networks.

with scango you can:
- Scan IP addresses on networks
- Scan ports on networks
- Scan IP addresses and ports on networks with a specific protocol`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
