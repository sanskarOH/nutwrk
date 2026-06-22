/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/sanskarOH/nutwrk/internal/test"

	"github.com/spf13/cobra"
)

var size string

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Measure download speed using a configurable test file",
	Long: `Run an internet download speed benchmark test by downloading
a test file from a remote server and measuring the transfer rate.

The command reports:
  • File size
  • Download progress
  • Total transfer time
  • Average download speed (Mbps)

Supported test sizes:
  • 1mb
  • 10mb
  • 100mb
  • 1gb

Examples:
  nutwrk test
  nutwrk test -s 10mb
  nutwrk test -s 100mb
  nutwrk test -s 1gb`,
	Run: func(cmd *cobra.Command, args []string) {
		test.Check(size)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().StringVarP(
		&size,
		"size",
		"s",
		"100mb",
		"size of the download test file (1mb, 10mb, 100mb, 1gb)",
	)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
