package main

import (
	"fmt"
	"os"

	"github.com/taxio/lostat/cmd"
)

var (
	version = "dev"
)

func main() {
	rootCmd := cmd.Root(version)
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
