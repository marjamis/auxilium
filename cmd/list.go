package cmd

import (
	"github.com/spf13/cobra"

	"github.com/marjamis/auxilium/pkg/engine"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists out all the steps in the supplied configuration file.",
	Run: func(cmd *cobra.Command, args []string) {
		engine.ListSteps(bbd.Steps)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
