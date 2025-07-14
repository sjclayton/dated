package format

import (
	"strings"
	"time"
)

func ClockFormat(t time.Time, clockFormat bool) string {
	switch clockFormat {
	case true:
		return t.Format("15:04")
	default:
		return t.Format("3:04PM")
	}
}

func TwelveHourTimeToWords(hours, minutes int, useOptional bool) string {
	if hours < 0 || hours > 23 || minutes < 0 || minutes > 59 {
		return "Time out of range"
	}

	// Convert hours to 12-hour format
	hour12 := hours % 12
	if hour12 == 0 {
		hour12 = 12 // Handle midnight and noon
	}
	hourWords := NumberToWords(hour12)

	// Handle the optional style
	if useOptional {
		if minutes == 0 {
			return hourWords + " O'Clock"
		} else if minutes == 15 {
			return "Quarter after " + hourWords
		} else if minutes == 30 {
			return "Half Past " + hourWords
		} else if minutes == 45 {
			// Calculate next hour in 12-hour format
			nextHour := (hours + 1) % 12
			if nextHour == 0 {
				nextHour = 12
			}
			nextHourWords := NumberToWords(nextHour)
			return "Quarter to " + nextHourWords
		} else if minutes < 30 {
			minuteWords := NumberToWords(minutes)
			return minuteWords + " after " + hourWords
		} else {
			// minutes > 30
			minsToNext := 60 - minutes
			minuteWords := NumberToWords(minsToNext)
			nextHour := (hours + 1) % 12
			if nextHour == 0 {
				nextHour = 12
			}
			nextHourWords := NumberToWords(nextHour)
			return minuteWords + " to " + nextHourWords
		}
	} else {
		// Primary style
		if minutes == 0 {
			return hourWords + " O'Clock"
		} else if minutes < 10 {
			minuteWords := NumberToWords(minutes)
			return hourWords + " O'" + minuteWords
		} else {
			minuteWords := NumberToWords(minutes)
			return hourWords + " " + minuteWords
		}
	}
}

func MilitaryTimeToWords(hours, minutes int) string {
	if hours < 0 || hours > 23 || minutes < 0 || minutes > 59 {
		return "Time out of range"
	}

	// Handle special case for midnight
	if hours == 0 && minutes == 0 {
		return "Zero Hundred"
	}

	var result string

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
