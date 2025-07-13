package cmd

import (
	"github.com/spf13/cobra"
)

var (
	hideMeridiem bool
	shortWkDay   bool
	shortMonth   bool
	showTime     bool
	timeFormat   string
	words        bool
)

func sharedFlagsBase(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&shortWkDay, "short-day", "D", false, "Use short weekday (Mon / Tue / Wed) <default: false>")
	cmd.Flags().BoolVarP(&shortMonth, "short-month", "M", false, "Use short month (Jan / Feb / Mar) <default: false>")
}

func sharedFlagsExtended(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&words, "words", "W", false, "Output words (example: 'Seven Fifteen' instead of '07:15', or 'January Thirtieth' instead of '01-30')")
}

func sharedFlagsTime(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&showTime, "time", "t", false, "Include time <default: false>")
	cmd.Flags().StringVarP(&timeFormat, "time-format", "T", "12h", "Time format: 12h/24h, <default: 12h>")
	cmd.Flags().BoolVarP(&hideMeridiem, "meridiem", "m", false, "Hide meridiem (AM/PM) <default: false>")
}
