package main

import (
	"fmt"
	"strings"
	"time"
)

func daySuffix(day int) string {
	switch day {
	case 1, 21, 31:
		return "st"
	case 2, 22:
		return "nd"
	case 3, 23:
		return "rd"
	default:
		return "th"
	}
}

func main() {
	t := time.Now()
	day := t.Day()
	hour := t.Hour()
	if hour == 0 {
		hour = 12 // Midnight
	} else if hour > 12 {
		hour -= 12 // Afternoon
	}
	formatted := fmt.Sprintf("%s %d%s %d:%02d%s",
		t.Format("Jan"),
		day,
		daySuffix(day),
		hour,
		t.Minute(),
		t.Format("PM"))
	fmt.Println(strings.ToLower(formatted))
}
