/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nutwrk",
	Short: "Network diagnostics and performance testing CLI",
	Long: `Nutwrk is a lightweight networking toolkit written in Go.

It provides utilities for measuring network latency,
testing connectivity, and benchmarking network performance
through an intuitive command-line interface.

Current features:
  • TCP-based ping
  • DNS resolution
  • Packet loss tracking
  • Custom ports
  • Latency statistics

Future features:
  • Download speed testing
  • Upload speed testing
  • ICMP ping
  • Network history
  • Server benchmarking`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.speedT.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
