package main

import (
	"scango/cmd"
	_ "scango/cmd/host"
	_ "scango/cmd/host/range_host"
	_ "scango/cmd/host/specific_host"
	_ "scango/cmd/port"
	_ "scango/cmd/port/range_port"
	_ "scango/cmd/port/specific_port"
)

func main() {
	cmd.Execute()
}
