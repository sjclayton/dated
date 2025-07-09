package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/sjclayton/dated/format"
)

func init() {
	RootCmd.AddCommand(nowCmd)
}

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Prints the current date and time nicely formatted (default)",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		day := now.Day()
		hour := now.Hour()
		if hour == 0 {
			hour = 12 // Midnight
		} else if hour > 12 {
			hour -= 12 // Afternoon
		}
		output := fmt.Sprintf("%s %d%s %d:%02d%s",
			now.Format("Jan"),
			day,
			format.NumberSuffix(day),
			hour,
			now.Minute(),
			now.Format("PM"))

		switch outputCase {
		case "upper":
			output = strings.ToUpper(output)
		case "lower":
			output = strings.ToLower(output)
		}
		fmt.Println(output)
	},
}
