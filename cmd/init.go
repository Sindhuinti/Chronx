package cmd

import (

	"github.com/Sindhuinti/chronx/pkg"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)




var initCmd = &cobra.Command{
	Use:                   "init",
	Short:                 "Add user's google calendar oauth token",
	Long:                  `Chronx init command is used to set up user's google account oauth to get google calendar token`,
	Args:                  cobra.ExactArgs(0),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		_,err := pkg.GetClient()
		if err!=nil{
			color.Red(err.Error())
			return
		}
		color.Green("Oauth Success")

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
