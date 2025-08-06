package main

import (
	"fmt"
	"os"

	"github.com/arian-press2015/uniac/internal/cli"
)

func main() {
	if err := cli.RunCLI(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
