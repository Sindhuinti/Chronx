package cmd

import (
	"github.com/Sindhuinti/chronx/pkg"
	"github.com/spf13/cobra"
)

var list string


var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get list of events from user's calendar",
	Long: `This command is used to list the avaliable 
events of the user's google calendar`,
	Run: func(cmd *cobra.Command, args []string) {
		if list==""{

			pkg.GetEvents()

		}else{
			pkg.GetOneEvent(list)
		}
	},
	Example: `chronx list
chronx list -l "Visit park`,
}

func init() {
	
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringVarP(&list, "list", "l", "", "view detailed info of a event")

}
