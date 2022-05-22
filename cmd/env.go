package cmd

import (
	"gits/impl"

	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Show env",
	Long:  `Show env`,
	Run: func(cmd *cobra.Command, args []string) {
		impl.Env(show_env, repo)
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
	envCmd.PersistentFlags().BoolVarP(&show_env, "show", "s", true, "Show current env.")
	envCmd.PersistentFlags().StringArrayVarP(&repo, "add", "a", []string{}, "Add a repo.")
}
