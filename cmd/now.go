package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/sjclayton/dated/format"
)

func init() {
	RootCmd.AddCommand(nowCmd)

	flagsBase(nowCmd)
}

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Prints the current date and time nicely formatted (default)",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()

		output := fmt.Sprintf("%s %d%s %s",
			now.Format("Jan"),
			now.Day(),
			format.NumberSuffix(now.Day()),
			format.ClockFormat(now, clockFormat),
		)

		output = format.TransformCase(output, outputCase)
		fmt.Println(output)
	},
}
