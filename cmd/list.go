package cmd

import (
	"github.com/Sindhuinti/chronx/pkg"
	"github.com/spf13/cobra"
)


var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get list of events from user's calendar",
	Long: `This command is used to list the avaliable 
events of the user's google calendar`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.GetEvents()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
