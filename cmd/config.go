/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config [set|get|update] [<option>=<value>] [--file <config_file>]",
	Short: "View or modify the Tracee Daemon configuration at runtime.",
	Long:  `View or modify the Tracee Daemon configuration at runtime.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
