package main

import (
	"flag"
	"fmt"
	"log"
)

const (
	Version = "0.2"
)

func main() {
	var confFile string
	var versionFlag bool
	flag.StringVar(&confFile, "conf", "config.yml", "config file path")
	flag.BoolVar(&versionFlag, "version", false, "show version")
	flag.BoolVar(&versionFlag, "v", false, "show version")
	flag.Parse()

	if versionFlag {
		fmt.Println(Version)
		return
	}

	conf, err := ParseConf(confFile)
	if err != nil {
		panic(err)
	}

	server := NewServer(conf)
	log.Fatal(server.Run())
}
