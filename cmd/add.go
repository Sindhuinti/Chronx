package cmd

import (

	"github.com/Sindhuinti/chronx/pkg"
	"github.com/spf13/cobra"
)

var (
	title       string
	description string
	colorId     string
	start       string
	end         string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new event to user's calendar",
	Long: `Chronx add command is used to create new events in google calendar
	by passing arguments like Title, description, start time, end time, event links etc`,
	Run: func(cmd *cobra.Command, args []string) {
		
		pkg.AddEvent(title,description,colorId,start,end)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().StringVarP(&title, "title", "t", "", "add title to event")
	addCmd.PersistentFlags().StringVarP(&description, "desc", "d", "", "add description to event")
	addCmd.PersistentFlags().StringVarP(&colorId, "id", "i", "", "add color id to event")
	addCmd.PersistentFlags().StringVarP(&start, "start", "s", "", "add start time to event")
	addCmd.PersistentFlags().StringVarP(&end, "end", "e", "", "add end time to event")

}
