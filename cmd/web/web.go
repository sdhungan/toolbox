/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package web

import (
	"github.com/spf13/cobra"
)

// webCmd represents the web command
var WebCmd = &cobra.Command{
	Use:   "web",
	Short: "This command Pallet contains commands regarding the webservice",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
