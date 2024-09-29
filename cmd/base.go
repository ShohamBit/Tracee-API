package cmd

import "github.com/spf13/cobra"

var baseCmd = &cobra.Command{
	Use:   "base",
	Short: "short about the use",
	Long:  "long about the use",
	Run: func(cmd *cobra.Command, args []string) {
		base(cmd, args)
	},
}

func base(cmd *cobra.Command, _ []string) {

}
