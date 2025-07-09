package format

import "strings"

func MilitaryTimeToWords(hours, minutes int) string {
	if hours < 0 || hours > 23 || minutes < 0 || minutes > 59 {
		return "Time out of range"
	}

	// Handle special case for midnight
	if hours == 0 && minutes == 0 {
		return "Zero Hundred"
	}

	result := ""

	// Convert hours
	if hours > 0 && hours < 10 {
		result = "Zero " + NumberToWords(hours)
	} else {
		result = NumberToWords(hours)
	}

	// Add "Hundred" only on the hour
	if minutes == 0 {
		result += " Hundred"
	}

	// Add minutes if non-zero
	if minutes > 0 {
		if hours > 0 && minutes < 10 {
			result += " O'" + NumberToWords(minutes)
		} else {
			result += " " + NumberToWords(minutes)
		}
	}

	return strings.TrimSpace(result)
}
