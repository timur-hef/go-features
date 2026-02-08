package iteration

import "strings"

func Repeat(symbol string) string {
	var result string

	for i := 0; i < 5; i++ {
		result = result + symbol
	}

	return result
}

// since string in go are immutable the concatenation will need memory allocation.
// strint builder can help with that
func RepeatOptimized(symbol string) string {
	var result strings.Builder

	for i := 0; i < 5; i++ {
		result.WriteString(symbol)
	}

	return result.String()
}
