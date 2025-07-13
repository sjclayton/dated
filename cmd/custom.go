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
	sharedFlagsBase(customCmd)

	// Define format flag with default value
	customCmd.Flags().StringVarP(&formatString, "format", "f", "<DW> the <D> of <M>, <Y> at <T>", "Custom format string")
}

var customCmd = &cobra.Command{
	Use:   "custom",
	Short: "Prints a custom formatted output string",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		output := formatCustom(now, formatString, shortWkDay, shortMonth, suffix, words)
		fmt.Println(format.TransformCase(output, outputCase))
	},
}

// formatCustom generates a formatted string based on the provided tags.
func formatCustom(t time.Time, formatStr string, shortWkDay, shortMonth, useSuffix, useWords bool) string {
	// Map placeholders to time formats
	replacements := map[string]string{
		"<DW>": t.Format("Monday"),
		"<M>":  t.Format("January"),
		"<Y>":  t.Format("2006"),
		"<T>":  fmt.Sprintf("%02d:%02d%s", t.Hour(), t.Minute(), t.Format("PM")),
		"<D>":  fmt.Sprintf("%d", t.Day()),
		"<DY>": strconv.Itoa(t.YearDay()),
	}

	if shortWkDay {
		replacements["<DW>"] = t.Format("Mon")
	}

	if shortMonth {
		replacements["<M>"] = t.Format("Jan")
	}

	switch {
	case useSuffix && useWords:
		replacements["<D>"] = format.NumberToWords(t.Day(), true)
		replacements["<DY>"] = format.NumberToWords(t.YearDay(), true)
		replacements["<Y>"] = format.YearToWords(t.Year(), false)
	case useSuffix:
		replacements["<D>"] = fmt.Sprintf("%d%s", t.Day(), format.NumberSuffix(t.Day()))
		replacements["<DY>"] = fmt.Sprintf("%d%s", t.YearDay(), format.NumberSuffix(t.YearDay()))
	case useWords:
		replacements["<D>"] = format.NumberToWords(t.Day(), false)
		replacements["<DY>"] = format.NumberToWords(t.YearDay(), false)
		replacements["<Y>"] = format.YearToWords(t.Year(), false)
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
