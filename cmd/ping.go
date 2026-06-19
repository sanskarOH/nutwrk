/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/sanskarOH/speedT/internal/ping"
	"github.com/spf13/cobra"
)

var count int
var port int
var timeout int

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ltc",
	Short: "Measure TCP latency to a host",
	Long: `
	Ping a host by establishing TCP connections to a target port and 
	measuring the connection latency.

	The command resolves the hostname, extracts an IPv4 address,
	attempts TCP connections, and reports latency statistics such as
	average response time and packet loss.

	Examples:
  	nutwrk ltc google.com
  	nutwrk ltc google.com -c 10
  	nutwrk ltc localhost -p 8080
	nutwrk ltc google.com -c 5 -p 8080 -t 5`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		host := args[0]

		err := ping.Run(host, count, strconv.Itoa(port), timeout)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
	pingCmd.Flags().IntVarP(
		&count,
		"count",
		"c",
		4,
		"number of ping attempts",
	)
	pingCmd.Flags().IntVarP(
		&port,
		"port",
		"p",
		443,
		"port of the server",
	)
	pingCmd.Flags().IntVarP(
		&timeout,
		"timeout",
		"t",
		1,
		"timeout in seconds",
	)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
