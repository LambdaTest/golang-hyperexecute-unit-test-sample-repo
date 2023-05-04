package case1

import "strings"

func checkCase(s string) string {
	if s == strings.ToUpper(s) {
		return "uppercase"
	} else if s == strings.ToLower(s) {
		return "lowercase"
	} else if strings.ContainsAny(s, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") && strings.ContainsAny(s, "abcdefghijklmnopqrstuvwxyz") {
		return "camelcase"
	} else {
		return "none"
	}
}