package main

import (
	"fmt"
	"os"

	"github.com/taxio/lostat/cmd"
)

const (
	version = "v0.0.1"
)

func main() {
	rootCmd := cmd.Root(version)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
