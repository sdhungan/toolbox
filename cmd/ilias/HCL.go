/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package ilias

import (
	"github.com/spf13/cobra"
)

// HclTempCmd represents the hclTemp command
var HclTempCmd = &cobra.Command{
	Use:   "hcl",
	Short: "This pallet contains function to generate templates utilising the excel and hcl file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
