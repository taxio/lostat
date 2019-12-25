package cmd

import "github.com/spf13/cobra"

// Root returns lostat cli root object
func Root(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "lostat",
		Short:   "Local repository status checker",
		Long:    "Local repository status checker",
		Version: version,
	}

	subCmds := []*cobra.Command{
		checkSubCmd(),
	}
	rootCmd.AddCommand(subCmds...)

	return rootCmd
}
