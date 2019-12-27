package main

import (
	"fmt"
	"os"

	"github.com/taxio/lostat/cmd"
)

var (
	version = "v0.0.3"
)

func main() {
	rootCmd := cmd.Root(version)
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
