package pbt

import "strings"

var ArabicToRoman = map[uint16]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var SortedRomanNumbers = []uint16{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func ConvertToRoman(arabic uint16) string {

	var result strings.Builder

	for arabic > 0 {
		var threshold uint16

		for _, num := range SortedRomanNumbers {
			if num <= arabic {
				threshold = num
				break
			}
		}

		arabic -= threshold
		result.WriteString(ArabicToRoman[threshold])
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	var result uint16
	var romanNumber string
	var arabicNum uint16

	for len(roman) > 0 {
		isInvalid := true

		for _, num := range SortedRomanNumbers {
			romanNumber = ArabicToRoman[num]

			if strings.HasPrefix(roman, romanNumber) {
				arabicNum = num
				isInvalid = false
				break
			}
		}

		if isInvalid {
			panic("invalid input")
		}

		roman = roman[len(romanNumber):]
		result += arabicNum
	}

	return result
}
