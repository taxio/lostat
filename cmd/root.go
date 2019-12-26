package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/taxio/lostat/checker"
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
			for _, repoPath := range args {
				chkr, err := checker.New(repoPath)
				if err != nil {
					return fmt.Errorf("%w", err)
				}
				hasChanges, err := chkr.HasChanges()
				if err != nil {
					return fmt.Errorf("%w", err)
				}
				if hasChanges {
					fmt.Println(repoPath)
				}
			}
			return nil
		},
	}

	return rootCmd
}
