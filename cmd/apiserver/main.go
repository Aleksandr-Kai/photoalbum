package main

import (
	"flag"
	"log"
	"photoalbum/internal/app/apiserver"

	"github.com/Aleksandr-Kai/logger"
	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./configs/apiserver.toml", "path to config file")
}

func main() {
	logger.LogToConsole(logger.Info, "Start main programm")
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
