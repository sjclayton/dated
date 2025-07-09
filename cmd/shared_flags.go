package cmd

import (
	"github.com/spf13/cobra"
)

func defineSharedFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&suffix, "suffix", "s", false, "Use ordinal suffix for dates (1st / 2nd / First / Second) <default: false>")
	cmd.Flags().BoolVarP(&words, "words", "W", false, "Output words (example: 'Seven Fifteen' instead of '07:15', or 'January Thirtieth' instead of '01-30')")
}
