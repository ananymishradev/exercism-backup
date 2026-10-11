package luhn

import (
	"strings"
)
func Valid(id string) bool {
	trimStr := strings.ReplaceAll(id, " ", "")
	if len(trimStr) <= 1 {
		return false
	}
	sum := 0
	isSecond := false

	for i := len(trimStr) - 1; i >= 0; i-- {
		char := trimStr[i]
		if char < '0' || char > '9' {
			return false
		}
		digit := int(char - '0')
		if isSecond {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		isSecond = !isSecond
	}
	return sum%10 == 0
}