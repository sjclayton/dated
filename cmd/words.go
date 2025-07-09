package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/sjclayton/dated/format"
)

func init() {
	RootCmd.AddCommand(wordsCmd)
}

var wordsCmd = &cobra.Command{
	Use:   "words",
	Short: "Prints the current date with year in words",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		month := now.Format("January")
		day := format.NumberToWords(now.Day(), true)
		year := format.YearToWords(now.Year())
		output := fmt.Sprintf("%s %s, %s", month, day, year)

		switch outputCase {
		case "upper":
			output = strings.ToUpper(output)
		case "lower":
			output = strings.ToLower(output)
		}
		fmt.Println(output)
	},
}
