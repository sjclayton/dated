package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/sjclayton/dated/format"
)

func init() {
	RootCmd.AddCommand(shortCmd)

	flagsBase(shortCmd)
	flagsSharedPreset(shortCmd)
}

var shortCmd = &cobra.Command{
	Use:   "short",
	Short: "Prints the date in short format (e.g., '01/30') with optional flags",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		month := fmt.Sprintf("%02d", int(now.Month()))
		day := fmt.Sprintf("%02d", now.Day())

		// Apply locale region formatting (NA: Month Day, EU: Day Month)
		var output string
		switch locale {
		case "EU":
			output = fmt.Sprintf("%s/%s", day, month)
		default: // NA or any other value
			output = fmt.Sprintf("%s/%s", month, day)
		}

		// Optionally add day of week
		if showDay {
			weekday := now.Format("Mon")
			output = fmt.Sprintf("%s %s", weekday, output)
		}

		// Optionally add year
		if showYear {
			year := now.Format("06")
			output = fmt.Sprintf("%s/%s", output, year)
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
