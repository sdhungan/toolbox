/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/sdhungan/toolbox/cmd/info"
	"github.com/sdhungan/toolbox/cmd/net"
	"github.com/sdhungan/toolbox/cmd/web"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "toolbox",
	Short: "This CLI tool contains various pallets of commands",
	Long: `
This is a CLI tool which contains various pallets of commands such as network, information and more.
To access the pallets use the following commands:

toolbox net [-h] - for network info
toolbox info [-h] - for information commands
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommandPallets() {
	rootCmd.AddCommand(net.NetCmd)
	rootCmd.AddCommand(info.InfoCmd)
	rootCmd.AddCommand(web.WebCmd)

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.toolbox.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommandPallets()

}
