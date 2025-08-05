package main

import (
	"fmt"
	"os"

	"github.com/arian-press2015/uniac/internal/loader"
)

func main() {
	// use viper to parse cmd args
	// call code from `pkg/core`

	loader, err := loader.NewLoader("sample.yaml")
	if err != nil {
		fmt.Println("Error creating config loader:", err)
		os.Exit(1)
	}

	world, err := loader.Load("sample.yaml")
	if err != nil {
		fmt.Println("Error loading and validating:", err)
		os.Exit(1)
	}

	fmt.Println("World parsed and validated successfully")

	fmt.Print(world)
}
