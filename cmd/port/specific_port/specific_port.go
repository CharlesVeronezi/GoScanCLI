package specific_port

import (
	"fmt"
	"net"
	"scango/cmd/port"
	"time"

	"github.com/spf13/cobra"
)

var (
	ip        string
	entryPort int
)

var SpecificPortCmd = &cobra.Command{
	Use:   "specific",
	Short: "Find out if a port is open on the host",
	Long: `Find out if a port is open on the host.

- Use the --port command to specify a different port.
- Use the --ip command to specify a different IP address.`,

	Run: func(cmd *cobra.Command, args []string) {
		err := FindSpecificPort(ip, entryPort)
		if err != nil {
			fmt.Println("Port ", entryPort, " is closed")
		} else {
			fmt.Println("Port ", entryPort, " is open")
		}
	},
}

func FindSpecificPort(ip string, entryPort int) error {
	address := fmt.Sprintf("%s:%d", ip, entryPort)
	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}

func init() {
	SpecificPortCmd.Flags().StringVarP(&ip, "ip", "i", "", "IP address of the host (format should be: 999.999.9.999)")
	SpecificPortCmd.Flags().IntVarP(&entryPort, "port", "p", 0, "Port number to check")
	SpecificPortCmd.MarkFlagRequired("ip")
	SpecificPortCmd.MarkFlagRequired("port")

	port.PortCmd.AddCommand(SpecificPortCmd)
}
