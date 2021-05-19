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

	markdownContent := resp.ToMarkdown(cfg.Sections)

	if cfg.Github == nil {
		fmt.Println(markdownContent)
	} else {
		gc := newGithubClient(cfg.Github)

		url, err := gc.uploadGist(markdownContent)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Gist created: %s\n", url)
	}
}
