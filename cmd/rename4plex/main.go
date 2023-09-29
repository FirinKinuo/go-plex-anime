package main

import (
	"github.com/FirinKinuo/rename4plex/cli"
	"github.com/FirinKinuo/rename4plex/config"
	"github.com/charmbracelet/log"
)

var (
	version    = "unknown"
	configPath = ""
)

func main() {
	var err error
	if configPath == "" {
		configPath, err = config.GetConfigFilePath()
		if err != nil {
			log.Fatal("get config path", "err", err)
		}
	}

	cfg := config.NewConfig(configPath)

	r4p := cli.NewRename4Plex(version, cfg)
	_ = r4p.Execute()
}
