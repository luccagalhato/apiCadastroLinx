package main

import (
	"apiCadastro/config"
	_ "embed"
	"flag"
	"log"
	"os"
)

func initConfig() {
	var createConfig bool
	flag.BoolVar(&createConfig, "c", false, "create config.yaml file")
	flag.Parse()
	if createConfig {
		if err := config.NewYaml("config.yaml"); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
}

func main() {
	initConfig()
}
