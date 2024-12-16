/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Returns the IP address of the current machine",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				ip4 := ipNet.IP.To4()
				ip4mask := net.IP(ipNet.Mask).To4()
				if ip4 != nil && ip4mask != nil {
					fmt.Println("IP address:", ip4.String())
					fmt.Println("Subnet mask:", ip4mask.String())
					return
				}
			}
		}

		fmt.Println("No IP address found")
	},
}

func init() {

	NetCmd.AddCommand(ipCmd)
}
