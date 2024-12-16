/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	urlPath string
	client  = http.Client{
		Timeout: time.Second * 2,
	}
)

func ping(domain string) (int, error) {
	url := "http://" + domain
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This pings a remote URL to check if it is reachable",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// LOGIC of the ping command here
		if len(args) == 1 {
			urlPath = args[0]
		} else if urlPath == "" {
			fmt.Println("Error: You must provide either one argument or use the --url flag.")
			return
		}

		resp, err := ping(urlPath)
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Printf("Respons of %v: %v", urlPath, resp)
		}

	},
}

func init() {

	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url to ping")

	// Here you will define your flags and configuration settings.
	NetCmd.AddCommand(pingCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
