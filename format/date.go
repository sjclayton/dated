package format

import "strings"

func YearToWords(year int, long ...bool) string {
	wantLong := false
	if len(long) > 0 {
		wantLong = long[0]
	}

	if year < 1970 || year > 9999 {
		return "Year out of range"
	}

	thousands := year / 1000
	hundreds := (year % 1000) / 100
	remainder := year % 100

	var result string

	// Handle even thousands (2000, 3000, etc.)
	if year%1000 == 0 {
		return ones[thousands] + " Thousand"
	}

	// Handle 1970-1999
	if year < 2000 {
		if wantLong {
			result = NumberToWords(year/100) + " Hundred"
		} else {
			result = NumberToWords(year / 100)
		}
	} else {
		// Handle 2000-9999
		if wantLong {
			result = ones[thousands] + " Thousand"
			if hundreds != 0 {
				result += " " + NumberToWords(hundreds) + " Hundred"
			}
		} else {
			result = NumberToWords(year / 100)
		}
	}

	// Append remainder if non-zero
	if remainder != 0 {
		result += " "
		if wantLong {
			result += "and " + NumberToWords(remainder)
		} else if year >= 2000 {
			// For short form, use "O'" for single-digit remainders
			if remainder < 10 {
				result += "O'" + ones[remainder]
			} else {
				result += NumberToWords(remainder)
			}
		} else {
			result += NumberToWords(remainder)
		}
	}

	return strings.TrimSpace(result)
}
