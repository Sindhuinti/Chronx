package cmd

import (
	"fmt"

	"github.com/Sindhuinti/chronx/pkg"
	"github.com/spf13/cobra"
)


var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an event of user's calendar",
	Long: `Chronx update command is used to update an existing events in google calendar
by passing arguments like Title of event and start time, end time`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		title := args[0]
		var start string
		var end string
		var desc string
		var colorId string
		
		fmt.Println("Please enter the values you want to update (otherwise hit enter):")
		fmt.Println("(Start Time): ")
		fmt.Scan(&start)
		fmt.Println("(End Time): ")
		fmt.Scan(&end)
		fmt.Println("(Description): ")
		fmt.Scan(&desc)
		fmt.Println("(colorId): ")
		fmt.Scan(&colorId)
		pkg.UpdateEvent(title,start,end,desc,colorId)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	
}
