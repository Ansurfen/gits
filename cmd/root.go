package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gits",
	Short: "Gits enhances traditional git command.",
	Long:  `Gits is the abbreviation of 'git super', which enhances traditional git command.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
