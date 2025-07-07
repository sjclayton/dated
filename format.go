package main

var ones = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

var teens = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}

var tens = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

var ordinalOnes = []string{"", "First", "Second", "Third", "Fourth", "Fifth", "Sixth", "Seventh", "Eighth", "Ninth"}

var ordinalTeens = []string{"Tenth", "Eleventh", "Twelfth", "Thirteenth", "Fourteenth", "Fifteenth", "Sixteenth", "Seventeenth", "Eighteenth", "Nineteenth"}

var ordinalTens = []string{"", "", "Twentieth", "Thirtieth", "Fortieth", "Fiftieth", "Sixtieth", "Seventieth", "Eightieth", "Ninetieth"}

func DaySuffix(day int) string {
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

func NumberToWords(n int, suffix bool) string {
	if n == 0 {
		return "Zero"
	}
	if n < 10 {
		if suffix {
			return ordinalOnes[n]
		}
		return ones[n]
	}
	if n < 20 {
		if suffix {
			return ordinalTeens[n-10]
		}
		return teens[n-10]
	}
	if n < 100 {
		tensDigit := n / 10
		onesDigit := n % 10
		if onesDigit == 0 {
			if suffix {
				return ordinalTens[tensDigit]
			}
			return tens[tensDigit]
		}
		separator := "-"
		if suffix {
			return tens[tensDigit] + separator + ordinalOnes[onesDigit]
		}
		return tens[tensDigit] + separator + ones[onesDigit]
	}
	if n < 1000 {
		hundreds := n / 100
		remainder := n % 100
		base := ones[hundreds] + " Hundred"
		if suffix && remainder == 0 {
			return base + "th"
		}
		if remainder == 0 {
			return base
		}
		return base + " " + NumberToWords(remainder, suffix)
	}
	if n >= 1000 {
		thousands := n / 1000
		remainder := n % 1000
		base := ones[thousands] + " Thousand"
		if suffix && remainder == 0 {
			return base + "th"
		}
		if remainder == 0 {
			return base
		}
		return base + " " + NumberToWords(remainder, suffix)
	}
	return ""
}

func YearToWords(year int, long bool) string {
	if year < 1000 || year > 9999 {
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

	// Handle 1000-1999
	if year < 2000 {
		if year < 1100 {
			result = "One Thousand"
		} else {
			if long {
				result = NumberToWords(year/100, false) + " Hundred"
			} else {
				result = NumberToWords(year/100, false)
			}
		}
	} else {
		// Handle 2000-9999
		if long || (hundreds == 0 && remainder < 10) {
			result = ones[thousands] + " Thousand"
			if hundreds != 0 {
				result += " " + NumberToWords(hundreds, false) + " Hundred"
			}
		} else {
			result = NumberToWords(year/100, false)
		}
	}

	// Append remainder if non-zero
	if remainder != 0 {
		result += " "
		if long && remainder < 10 {
			result += "and "
		} else if remainder < 10 && year >= 1100 && year < 2000 {
			result += "O' "
		}
		result += NumberToWords(remainder, false)
	}

	return result
}
