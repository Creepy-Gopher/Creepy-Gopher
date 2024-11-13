package utils

import (
	"strings"
)

func ReplacePersianDigits(input string) string {
	persianDigitsMap := map[rune]rune{
		'٠': '0',
		'۰': '0',
		'١': '1',
		'۱': '1',
		'٢': '2',
		'۲': '2',
		'٣': '3',
		'۳': '3',
		'۴': '4',
		'٤': '4',
		'۵': '5',
		'٥': '5',
		'۶': '6',
		'٦': '6',
		'٧': '7',
		'۷': '7',
		'٨': '8',
		'۸': '8',
		'٩': '9',
		'۹': '9',
	}

	var builder strings.Builder
	builder.Grow(len(input))

	for _, char := range input {
		englishDigit, ok := persianDigitsMap[char]
		if ok {
			builder.WriteRune(englishDigit)
		} else {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}
