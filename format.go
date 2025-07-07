package main

import "strings"

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

func NumberToWords(n int, suffix ...bool) string {
	wantSuffix := false
	if len(suffix) > 0 {
		wantSuffix = suffix[0]
	}

	// Handle zero
	if n == 0 {
		if wantSuffix {
			return "Zeroth"
		}
		return "Zero"
	}

	// Handle negative numbers
	if n < 0 {
		return "Negative " + NumberToWords(-n)
	}

	// Handle numbers up to 99
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

	// Handle numbers up to 999 (explicit hundreds)
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
		if wantSuffix {
			return base + " " + NumberToWords(remainder, true)
		}
		return base + " " + NumberToWords(remainder)
	}

	// Handle larger numbers
	result := ""
	num := n
	scaleIndex := 0

	for num > 0 && scaleIndex < len(scales) {
		chunk := num % 1000
		if chunk > 0 {
			isLastChunk := num < 1000
			chunkStr := NumberToWords(chunk, wantSuffix && !isLastChunk)
			if result != "" {
				result = chunkStr + " " + scales[scaleIndex] + " " + result
			} else {
				result = chunkStr + " " + scales[scaleIndex]
			}
		}
		num /= 1000
		scaleIndex++
	}

	return strings.TrimSpace(result)
}

func YearToWords(year int, long bool) string {
	if year < 1970 || year > 9999 {
		return "Year out of range"
	}

	thousands := year / 1000
	hundreds := (year % 1000) / 100
	remainder := year % 100
	result := ""

	// Handle even thousands (2000, 3000, etc.)
	if year%1000 == 0 {
		return ones[thousands] + " Thousand"
	}

	// Handle 1970-1999
	if year < 2000 {
		if long {
			result = NumberToWords(year/100) + " Hundred"
		} else {
			result = NumberToWords(year / 100)
		}
	} else {
		// Handle 2000-9999
		if long || (hundreds == 0 && remainder < 10) {
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
		if long && remainder < 10 {
			result += "and "
		} else if remainder < 10 && year >= 1970 && year < 2000 {
			result += "O' "
		}
		result += NumberToWords(remainder)
	}

	return result
}
