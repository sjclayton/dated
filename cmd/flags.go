package cmd

import (
	"github.com/spf13/cobra"
)

var (
	clockFormat bool
	locale      string
	shortMonth  bool
	shortWkDay  bool
	shortYear   bool
	showDay     bool
	showTime    bool
	showYear    bool
	suffix      bool
	words       bool
)

// Common flags
func flagsBase(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&clockFormat, "24h", "T", false, "Use 24-hour time format")
}

// Flags that can be used with custom format strings
func flagsCustomFmt(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&shortWkDay, "short-day", "D", false, "Use short weekday (Mon / Tue / Wed)")
	cmd.Flags().BoolVarP(&shortMonth, "short-month", "M", false, "Use short month (Jan / Feb / Mar)")
	cmd.Flags().BoolVarP(&shortYear, "short-year", "Y", false, "Use short year ('25 / Twenty Twenty-Five)")
	cmd.Flags().BoolVarP(&suffix, "suffix", "s", false, "Use ordinal suffix for dates (1st / 2nd / First / Second)")
	cmd.Flags().BoolVarP(&words, "words", "W", false, "Output words (e.g. 'Seven Fifteen' instead of '7:15', or 'January Thirtieth' instead of '01/30')")
}

// Shared flags usable by customizable presets
func flagsSharedPreset(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&showDay, "day", "d", false, "Include day of week")
	cmd.Flags().StringVarP(&locale, "locale", "l", "NA", "Locale date format: NA/EU")
	cmd.Flags().BoolVarP(&showTime, "time", "t", false, "Include time")
	cmd.Flags().BoolVarP(&showYear, "year", "y", false, "Include year")
}
