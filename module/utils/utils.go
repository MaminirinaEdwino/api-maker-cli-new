package utils

import "fmt"

func AddComaOrNot(idx int, lenAttr int) string {
	if idx < lenAttr {
		return ","
	}
	return ""
}
func WriteErrorCheker(message string) string {
	return fmt.Sprintf("if err != nil {\nfmt.Println(\"%s\", err)\n}", message)
}