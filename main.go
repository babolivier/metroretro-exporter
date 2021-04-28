package main

import (
	"flag"
	"fmt"
)

var (
	configPath = flag.String("c", "config.yaml", "Path to the configuration file")
	boardID    = flag.String("b", "", "ID of the board to export")
)

func main() {
	flag.Parse()

	if *boardID == "" {
		panic("board ID is mandatory")
	}

	cfg, err := newConfig(*configPath)
	if err != nil {
		panic(err)
	}

	client, err := newClient(cfg.Metroretro)
	if err != nil {
		panic(err)
	}

	resp, err := client.getBoard(*boardID)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.ToMarkdown(cfg.Sections))
}
