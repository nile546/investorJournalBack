package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// Config ...
type Config struct {
	Address string
}

func getConf() {
	c := "./config/config.toml"
	var conf Config

	d, err := toml.DecodeFile(c, &conf)
	if err != nil {
		fmt.Println("Cant loading config", err)
		os.Exit(1)
	}

	fmt.Println(d)
}

func main() {
	getConf()
}
