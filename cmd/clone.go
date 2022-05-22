package cmd

import (
	"gits/impl"

	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Better clone",
	Long:  `Better clone`,
	Run: func(cmd *cobra.Command, args []string) {
		impl.GitClone(args)
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
