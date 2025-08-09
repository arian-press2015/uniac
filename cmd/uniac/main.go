package main

import (
	"fmt"
	"log"
	"os"

	"github.com/arian-press2015/uniac/internal/cli"
)

func main() {
	c, err := cli.NewCLI()
	if err != nil {
		log.Fatalf("CLI initialization failed: %v", err)
	}

	if err := c.RunCLI(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
