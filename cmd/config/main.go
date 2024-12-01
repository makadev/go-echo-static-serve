package main

import (
	"example/go-echo-stuff/webserver/internal/config"
	"flag"
	"fmt"
)

func main() {
	// cli command -dump for duming the config
	dumpConfig := flag.Bool("dump", false, "dump configuration")
	flag.Parse()

	// load config
	config := config.NewConfig()
	config.Load()

	if *dumpConfig {
		out := config.Dump()
		fmt.Print(out)
	}
}
