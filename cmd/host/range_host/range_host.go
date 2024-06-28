package range_host

import (
	"fmt"
	"net"
	"scango/cmd/host"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var (
	ip   string
	port int = 80
	min  int = 0
	max  int = 999
	open int = 0
)

var RangeHostCmd = &cobra.Command{
	Use:   "range",
	Short: "Discover if IPs in a range are active on the network.",
	Long: `Discover if IPs in a range are active on the network.
	
By default, the packet is sent to port 80. 
Use the --port command to specify a different port.
Use the --ip command to specify a base IP address (e.g., 192.168.5.).`,

	Run: func(cmd *cobra.Command, args []string) {

		if open != 0 && open != 1 {
			panic("invalid value for --open flag: must be 0 or 1")
		}

		var wg sync.WaitGroup

		for i := min; i <= max; i++ {
			fullIP := fmt.Sprintf("%s%d", ip, i)
			wg.Add(1)
			go func(ip string) {
				defer wg.Done()
				err := FindRangeHost(fullIP, port)
				if err != nil {
					if open == 1 {
						fmt.Printf("Host %s is not active\n", fullIP)
					}
				} else {
					fmt.Printf("Host %s is active\n", fullIP)
				}
			}(fullIP)
		}

		wg.Wait()
	},
}

func FindRangeHost(ip string, port int) error {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		return err
	}
	conn.Close()

	return nil
}

func init() {
	RangeHostCmd.Flags().StringVarP(&ip, "ip", "i", "", "IP address of the host (format should be: 999.999.9.)")
	RangeHostCmd.Flags().IntVarP(&min, "beg", "b", 0, "The beginning of the IP range")
	RangeHostCmd.Flags().IntVarP(&max, "end", "e", 999, "The end of the IP range")
	RangeHostCmd.Flags().IntVarP(&port, "port", "p", 80, "Port number to check")
	RangeHostCmd.Flags().IntVarP(&open, "open", "o", 0, "Filter hosts: 1 for all hosts, 0 for active hosts only.")
	RangeHostCmd.MarkFlagRequired("ip")

	host.HostCmd.AddCommand(RangeHostCmd)
}
