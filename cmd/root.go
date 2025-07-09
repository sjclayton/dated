// Package cmd provides the all cmds valid for the application.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	outputCase string
	suffix     bool
	words      bool
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&outputCase, "case", "c", "", "Set output case: u/upper, l/lower")
}

var RootCmd = &cobra.Command{
	Use:   "dated",
	Short: "A simple CLI tool to output dates (and times) in various formats.",
	// Long:  `This application provides the ability to format date (and time) output in various ways.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Only run the default if no subcommand or flags like -h/--help are present
		if len(args) == 0 && !cmd.Flags().Changed("help") {
			nowCmd.Run(cmd, args)
		}
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
