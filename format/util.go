// Package format provides formatting functions used by the application.
package format

import (
	"strings"
)

// Lookup tables
var (
	ones         = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	ordinalOnes  = []string{"", "First", "Second", "Third", "Fourth", "Fifth", "Sixth", "Seventh", "Eighth", "Ninth"}
	teens        = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	ordinalTeens = []string{"Tenth", "Eleventh", "Twelfth", "Thirteenth", "Fourteenth", "Fifteenth", "Sixteenth", "Seventeenth", "Eighteenth", "Nineteenth"}
	tens         = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	ordinalTens  = []string{"", "", "Twentieth", "Thirtieth", "Fortieth", "Fiftieth", "Sixtieth", "Seventieth", "Eightieth", "Ninetieth"}
	scales       = []string{"", "Thousand", "Million", "Billion", "Trillion"}
)

// NumberSuffix returns a string representation of the ordinal suffix (e.g., "st", "nd", "rd", "th") for a given integer.
func NumberSuffix(n int) string {
	lastTwoDigits := n % 100
	lastDigit := n % 10

	// Numbers ending in 11, 12, 13 always get "th"
	if lastTwoDigits == 11 || lastTwoDigits == 12 || lastTwoDigits == 13 {
		return "th"
	}

	// Apply suffix based on last digit
	switch lastDigit {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

// NumberToWords converts a number to its word representation.
// If suffix (optional) is true, an ordinal suffix (e.g., "th", "st", "nd", "rd") is appended to the output.
//
// For example, NumberToWords(10) returns "Ten" and NumberToWords(10, true) returns "Tenth".
func NumberToWords(n int, suffix ...bool) string {
	wantSuffix := false
	if len(suffix) > 0 {
		wantSuffix = suffix[0]
	}

	// Handle numbers < 1000, including negative numbers and zero
	if n == 0 {
		if wantSuffix {
			return "Zeroth"
		}
		return "Zero"
	}

	if n < 0 {
		return "Negative " + NumberToWords(-n, wantSuffix)
	}

	if n < 10 {
		if wantSuffix {
			return ordinalOnes[n]
		}
		return ones[n]
	}

	if n < 20 {
		if wantSuffix {
			return ordinalTeens[n-10]
		}
		return teens[n-10]
	}

	if n < 100 {
		tensDigit := n / 10
		onesDigit := n % 10
		if onesDigit == 0 {
			if wantSuffix {
				return ordinalTens[tensDigit]
			}
			return tens[tensDigit]
		}
		separator := "-"
		if wantSuffix {
			return tens[tensDigit] + separator + ordinalOnes[onesDigit]
		}
		return tens[tensDigit] + separator + ones[onesDigit]
	}

	if n < 1000 {
		hundredsDigit := n / 100
		remainder := n % 100
		base := ones[hundredsDigit] + " Hundred"
		if remainder == 0 {
			if wantSuffix {
				return base + "th"
			}
			return base
		}
		return base + " " + NumberToWords(remainder, wantSuffix)
	}

	// Handle larger numbers
	var result string
	var scaleIndex int

	for n > 0 && scaleIndex < len(scales) {
		chunk := n % 1000
		if chunk > 0 {
			chunkStr := NumberToWords(chunk, false)
			scale := scales[scaleIndex]
			if result != "" {
				result = chunkStr + " " + scale + " " + result
			} else {
				result = chunkStr + " " + scale
			}
		}
		n /= 1000
		scaleIndex++
	}

	if wantSuffix {
		result = strings.TrimSpace(result) + NumberSuffix(n)
	}

	return strings.TrimSpace(result)
}

// TransformCase transforms the case of the output string if the case flag is passed with a valid option.
func TransformCase(output, outputCase string) string {
	switch outputCase {
	case "u", "upper":
		return strings.ToUpper(output)
	case "l", "lower":
		return strings.ToLower(output)
	default:
		return output
	}
}
