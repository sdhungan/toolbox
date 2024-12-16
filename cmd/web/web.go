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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
