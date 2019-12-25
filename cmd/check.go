package cmd

import (
	"github.com/spf13/cobra"
)

func checkSubCmd() *cobra.Command {
	checkCmd := &cobra.Command{
		Use:   "check",
		Short: "check repository status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return checkCmd
}
