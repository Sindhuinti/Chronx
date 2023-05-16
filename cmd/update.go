package cmd

import (
	"github.com/Sindhuinti/chronx/pkg"
	"github.com/spf13/cobra"
)


var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an event of user's calendar",
	Long: `Chronx update command is used to update an existing events in google calendar
by passing arguments like Title of event and start time, end time`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.UpdateEvent()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	
}
