package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/sjclayton/dated/format"
)

var formatString string

func init() {
	RootCmd.AddCommand(customCmd)

	flagsBase(customCmd)
	flagsCustomFmt(customCmd)

	// Define format flag with default value
	customCmd.Flags().StringVarP(&formatString, "format", "f", "<DW> <M> <D> <Y> @ <T>", "Custom format string")
}

var customCmd = &cobra.Command{
	Use:   "custom",
	Short: "Prints a custom formatted output string",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		output := formatCustom(now, formatString, shortWkDay, shortMonth, shortYear, suffix, words)
		fmt.Println(format.TransformCase(output, outputCase))
	},
}

// formatCustom generates a formatted string based on the provided tags.
func formatCustom(t time.Time, formatStr string, shortWkDay, shortMonth, shortYear, useSuffix, useWords bool) string {
	replacements := map[string]string{
		"<D>":  fmt.Sprintf("%d", t.Day()),
		"<DW>": t.Format("Monday"),
		"<DY>": strconv.Itoa(t.YearDay()),
		"<M>":  t.Format("January"),
		"<Y>":  t.Format("2006"),
		"<T>":  format.ClockFormat(t, clockFormat),
	}

	// Weekday
	if shortWkDay {
		replacements["<DW>"] = t.Format("Mon")
	}

	// Month
	if shortMonth {
		replacements["<M>"] = t.Format("Jan")
	}

	// Year
	switch {
	case shortYear && words:
		replacements["<Y>"] = format.YearToWords(t.Year(), false)
	case shortYear && !words:
		replacements["<Y>"] = "'" + t.Format("06")
	case !shortYear && words:
		replacements["<Y>"] = format.YearToWords(t.Year(), true)
	}

	// Day and YearDay
	switch {
	case useSuffix && useWords:
		replacements["<D>"] = format.NumberToWords(t.Day(), true)
		replacements["<DY>"] = format.NumberToWords(t.YearDay(), true)
		if !clockFormat {
			replacements["<T>"] = format.TwelveHourTimeToWords(t.Hour(), t.Minute(), false)
		} else {
			replacements["<T>"] = format.MilitaryTimeToWords(t.Hour(), t.Minute())
		}
	case useSuffix:
		replacements["<D>"] = fmt.Sprintf("%d%s", t.Day(), format.NumberSuffix(t.Day()))
		replacements["<DY>"] = fmt.Sprintf("%d%s", t.YearDay(), format.NumberSuffix(t.YearDay()))
	case useWords:
		replacements["<D>"] = format.NumberToWords(t.Day(), false)
		replacements["<DY>"] = format.NumberToWords(t.YearDay(), false)
		if !clockFormat {
			replacements["<T>"] = format.TwelveHourTimeToWords(t.Hour(), t.Minute(), false)
		} else {
			replacements["<T>"] = format.MilitaryTimeToWords(t.Hour(), t.Minute())
		}
	}

	return replacePlaceholders(formatStr, replacements)
}

// replacePlaceholders substitutes placeholders in the format string with their corresponding values.
func replacePlaceholders(formatStr string, replacements map[string]string) string {
	result := formatStr
	for placeholder, value := range replacements {
		result = strings.ReplaceAll(result, placeholder, value)
	}
	return result
}
