package main

import (
	"fmt"
	"os"

	"github.com/nile546/diplom/config"
)

func main() {

	c := config.NewConfig()

	err := c.LoadConfig()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//fmt.Println(c)
}
