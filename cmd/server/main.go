package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nile546/diplom/config"
	"github.com/nile546/diplom/internal/apiserver"
)

func main() {
	filePath := "./config/config.json"
	configfilePath := flag.String("config", filePath, "use config flag for set config json path")
	flag.Parse()
	c := config.NewConfig()

	err := c.LoadConfig(*configfilePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	api, err := apiserver.NewAPIServer(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = api.Start()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
