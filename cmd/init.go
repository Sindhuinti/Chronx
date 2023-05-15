package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Add user's google calendar oauth token",
	Long: `Chronx init command is used to set up user's google account oauth to get google calendar token`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

}
