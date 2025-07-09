package cmd

import (
	"fmt"
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
		output = format.TransformCase(output, outputCase)

		fmt.Println(output)
	},
}
