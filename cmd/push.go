package cmd

import (
	"gits/impl"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Better push",
	Long:  `Better push`,
	Run: func(cmd *cobra.Command, args []string) {
		impl.GitPush(isAll, redirectTarget, args...)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.PersistentFlags().BoolVarP(&isAll, "all", "a", true, "")
	pushCmd.PersistentFlags().BoolVarP(&time_to_run, "time", "t", true, "")
	pushCmd.PersistentFlags().StringVarP(&redirectTarget, "redirect", "r", "", "")
}
