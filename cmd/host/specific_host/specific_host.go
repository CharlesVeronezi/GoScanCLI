package specific_host

import (
	"fmt"
	"net"
	"scango/cmd/host"
	"time"

	"github.com/spf13/cobra"
)

var (
	ip   string
	port int = 80
)

var SpecificHostCmd = &cobra.Command{
	Use:   "specific",
	Short: "Discover if an IP is active on the network.",
	Long: `Discover if an IP is active on the network.

- By default, the packet is sent to port 80. 
- Use the --port command to specify a different port.
- Use the --ip command to specify a different IP address.`,

	Run: func(cmd *cobra.Command, args []string) {
		err := FindSpecificHost(ip, port)
		if err != nil {
			fmt.Println("Host " + ip + " is not active")
		} else {
			fmt.Println("Host " + ip + " is active")
		}
	},
}

func FindSpecificHost(ip string, port int) error {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		return err
	}
	conn.Close()

	return nil
}

func init() {
	SpecificHostCmd.Flags().StringVarP(&ip, "ip", "i", "", "IP address of the host (format should be: 999.999.9.999)")
	SpecificHostCmd.Flags().IntVarP(&port, "port", "p", 80, "Port number to check")
	SpecificHostCmd.MarkFlagRequired("ip")

	host.HostCmd.AddCommand(SpecificHostCmd)
}
