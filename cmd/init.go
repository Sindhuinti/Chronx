package cmd

import (
	"github.com/Sindhuinti/chronx/pkg"
	"github.com/spf13/cobra"
)


var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Add user's google calendar oauth token",
	Long: `Chronx init command is used to set up user's google account oauth to get google calendar token`,
	Args: cobra.ExactArgs(0),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.GetClient()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
