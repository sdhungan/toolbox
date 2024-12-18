/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package web

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sdhungan/toolbox/logger"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var startEchoCmd = &cobra.Command{
	Use:   "startCobra",
	Short: "Starts a web server on port 8080 in a new terminal utilising Echo pkg",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Launching the web server in a new terminal window...")
		log := logger.GetLogger()
		log.Info("Launch web server using new terminal.")
		// Define the command to launch a new terminal and run the web server
		var terminalCommand *exec.Cmd

		if runtime.GOOS == "windows" {
			/* On Windows: Open a new `cmd` window on the current directory
			"cmd /C start" 							will start a new terminal
			"start cmd /K"							will keep the terminal open after the command finishes (use /C to auto close after command finishes)
			"toolbox.exe web startweb-server"		Is the actual command run in the new terminal
			*/
			terminalCommand = exec.Command("cmd", "/C", "start", "cmd", "/K", "toolbox.exe web startCobra startCobra-server")
		} else {
			fmt.Println("This command is currently only supported on Windows.")
			return
		}

		// Start the new terminal process
		err := terminalCommand.Start()
		if err != nil {
			fmt.Printf("Failed to open new terminal: %s\n", err)
			return
		}

		fmt.Println("Web server is now running in a separate terminal.")
	},
}

// startWebCmd represents the startWeb command
var startEchoCmdServer = &cobra.Command{
	Use:   "c",
	Short: "Stars a web server on port 8080 in a new terminal utilising the basic HTTP go pkg",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.GetLogger()

		exePath, err := os.Executable()
		if err != nil {
			log.Error("Failed to get executable path ", zap.String("path", exePath), zap.Error(err))
			return
		}
		fmt.Println("Echo server is running on port 8080")

		e := echo.New()
		e.GET("/", GetRoot)

		e.Logger.Fatal(e.Start(":8080"))
	},
}

func init() {
	WebCmd.AddCommand(startEchoCmdServer)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startEchoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startEchoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
