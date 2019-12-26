package cmd

import (
	"fmt"
	"github.com/taxio/lostat/checker"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// Root returns lostat cli root object
func Root(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "lostat",
		Short:   "Local repository status checker",
		Long:    "Local repository status checker",
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("repository path not specified")
			}
			nParallel, err := cmd.Flags().GetInt("parallel")
			if err != nil {
				return fmt.Errorf("%w", err)
			}
			limit := make(chan bool, nParallel)
			for _, repoPath := range args {
				limit <- true
				go func(p string) {
					defer func() {
						<-limit
					}()
					chkr, err := checker.New(p)
					if err != nil {
						_, _ = fmt.Fprintf(os.Stderr, "%s: %s\n", p, err.Error())
						return
					}
					hasChanges, err := chkr.HasChanges()
					if err != nil {
						_, _ = fmt.Fprintf(os.Stderr, "%s: %s\n", p, err.Error())
						return
					}
					if hasChanges {
						fmt.Println(p)
					}
				}(repoPath)
			}
			return nil
		},
		SilenceErrors: true,
	}

	rootCmd.Flags().Int("parallel", runtime.NumCPU(), "number of worker process")

	return rootCmd
}
