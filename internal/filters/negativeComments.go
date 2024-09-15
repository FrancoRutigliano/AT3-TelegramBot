package filters

import "strings"

var negativeWords = []string{"estafadores", "ladrones", "chorros", "malos", "basura"}

func IsNegativeComment(message string) bool {
	for _, word := range negativeWords {
		if strings.Contains(strings.ToLower(message), word) {
			return true
		}
	}

	return false
}
