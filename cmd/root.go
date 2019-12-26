package cmd

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/spf13/cobra"
	"github.com/taxio/lostat/checker"
	"github.com/taxio/lostat/log"
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
			wg := &sync.WaitGroup{}
			limit := make(chan bool, nParallel)
			for _, repoPath := range args {
				limit <- true
				wg.Add(1)
				go func(p string) {
					defer func() {
						<-limit
						wg.Done()
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
						_, _ = fmt.Fprintln(os.Stdout, p)
					}
				}(repoPath)
			}
			wg.Wait()
			return nil
		},
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			verbose, err := cmd.PersistentFlags().GetBool("verbose")
			if err != nil {
				return fmt.Errorf("%w", err)
			}
			if verbose {
				log.SetVerboseLogger(os.Stdout)
				log.Println("verbose on")
			}
			return nil
		},
	}

	rootCmd.Flags().Int("parallel", runtime.NumCPU(), "number of worker process")
	rootCmd.PersistentFlags().Bool("verbose", false, "print log for developer")

	return rootCmd
}
