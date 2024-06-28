package range_port

import (
	"fmt"
	"net"
	"scango/cmd/port"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var (
	ip        string
	entryPort int = 0
	endPort   int = 65535
	open      int = 0
)

var RangePortCmd = &cobra.Command{
	Use:   "range",
	Short: "Discover if IPs in a range are active on the network.",
	Long:  `Discover if IPs in a range are active on the network.`,

	Run: func(cmd *cobra.Command, args []string) {

		if open != 0 && open != 1 {
			panic("invalid value for --open flag: must be 0 or 1")
		}

		var wg sync.WaitGroup
		wg.Add(endPort - entryPort + 1)

		for port := entryPort; port <= endPort; port++ {
			go func(port int) {
				defer wg.Done()
				err := FindRangePort(ip, port)
				if err != nil {
					if open == 1 {
						fmt.Printf("Port %d is closed\n", port)
					}
				} else {
					fmt.Printf("Port %d is open\n", port)
				}
			}(port)
		}

		wg.Wait()
	},
}

func FindRangePort(ip string, port int) error {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		return err
	}
	conn.Close()

	return nil
}

func init() {
	RangePortCmd.Flags().StringVarP(&ip, "ip", "i", "", "IP address of the host (format should be: 999.999.9.999)")
	RangePortCmd.Flags().IntVarP(&entryPort, "bP", "b", 0, "The beginning of the Port range")
	RangePortCmd.Flags().IntVarP(&endPort, "eP", "e", 65535, "The end of the Port range")
	RangePortCmd.Flags().IntVarP(&open, "open", "o", 0, "Filter ports: 1 for all ports, 0 for open ports only")
	RangePortCmd.MarkFlagRequired("ip")

	port.PortCmd.AddCommand(RangePortCmd)
}
