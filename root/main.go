package main

import (
	"fmt"
	"flag"
	"os"
	"account-management/root/server"
	"account-management/root/config"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	server.Init()
}
