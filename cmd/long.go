package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/sjclayton/dated/format"
)

func init() {
	RootCmd.AddCommand(longCmd)

	flagsBase(longCmd)
	flagsSharedPreset(longCmd)

	longCmd.Flags().BoolVarP(&suffix, "suffix", "s", false, "Use ordinal suffix for dates (1st / 2nd)")
}

var longCmd = &cobra.Command{
	Use:   "long",
	Short: "Prints the date in long format (e.g., 'January 15') with optional flags",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		month := now.Format("January")
		day := fmt.Sprintf("%d", now.Day())
		if suffix {
			day = fmt.Sprintf("%d%s", now.Day(), format.NumberSuffix(now.Day()))
		}

		// Apply locale region formatting (NA: Month Day, EU: Day Month)
		var output string
		switch locale {
		case "EU":
			output = fmt.Sprintf("%s %s", day, month)
		default: // NA or any other value
			output = fmt.Sprintf("%s %s", month, day)
		}

		// Optionally add day of week
		if showDay {
			weekday := now.Format("Monday")
			output = fmt.Sprintf("%s, %s", weekday, output)
		}

		// Optionally add year
		if showYear {
			year := now.Format("2006")
			output = fmt.Sprintf("%s %s", output, year)
		}

		// Optionally add time
		if showTime {
			clock := format.ClockFormat(now, clockFormat)
			output = fmt.Sprintf("%s %s", output, clock)
		}

		output = format.TransformCase(output, outputCase)
		fmt.Println(output)
	},
}
