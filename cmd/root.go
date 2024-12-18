/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/sdhungan/toolbox/cmd/ilias"
	"github.com/sdhungan/toolbox/cmd/info"
	"github.com/sdhungan/toolbox/cmd/net"
	"github.com/sdhungan/toolbox/cmd/web"
	"github.com/sdhungan/toolbox/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var log *zap.Logger

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
		log.Error("Command execute failed", zap.Error(err))
		os.Exit(1)
	}
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears the terminal screen",
	Long:  "Clears the terminal screen to make it easier to work with the CLI tool.",
	Run: func(cmd *cobra.Command, args []string) {
		switch runtime.GOOS {
		case "windows":
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
		default: // Unix-based systems
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	},
}

func init() {
	addSubCommandPallets()
	log = logger.GetLogger()

	// Assign the Run function after rootCmd is declared to avoid initialization cycle
	rootCmd.Run = func(cmd *cobra.Command, args []string) {

		log.Info("Interactive mode started...")
		fmt.Println("Welcome to the Toolbox interactive mode!")
		fmt.Println("Type '\x1b[33mhelp\x1b[0m' to see the list of commands or '\x1b[33mexit\x1b[0m' to quit the CLI")
		cmd.Help()

		startInteractiveMode()
	}
}

func startInteractiveMode() {
	// For the interactive mode we read in the input of the user using bufio.Scanner
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Infinite loop to keep the interactive mode running and waiting for user input
		fmt.Print("\ntoolbox> ")

		// Scan the input from the user until the user presses enter
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text()) // remove leading and trailing spaces

		if input == "" { // User presses enter without input text
			continue
		}

		if input == "exit" || input == "quit" { // User wants to exit the CLI
			log.Info("Exiting the toolbox CLI")
			break // Break infinite loop to exit the CLI
		}

		// If the user input is not empty and not exit, then we will execute the command
		args := strings.Split(input, " ") // Split the input into command and arguments by spaces
		rootCmd.SetArgs(args)             // Set the arguments for the root command as if the user typed them in the terminal

		if err := rootCmd.Execute(); err != nil {
			log.Error("Command execute failed: ", zap.Error(err))
		}

	}
}

func addSubCommandPallets() {
	rootCmd.AddCommand(net.NetCmd)
	rootCmd.AddCommand(info.InfoCmd)
	rootCmd.AddCommand(web.WebCmd)
	rootCmd.AddCommand(clearCmd)
	rootCmd.AddCommand((ilias.HclTempCmd))
	rootCmd.CompletionOptions.DisableDefaultCmd = true

}
