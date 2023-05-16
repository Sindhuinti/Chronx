package cmd

import (

	"github.com/Sindhuinti/chronx/pkg"
	"github.com/spf13/cobra"
)


var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing event in user's calendar",
	Long: `Chronx delete command is used to delete an existing event in 
google calendar by passing title as argument`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		pkg.DeleteEvent(title)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

}
