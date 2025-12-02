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

func ErrorChecker(err error) {
	if err != nil {
		panic(err)
	}
}

// Fonction qui scanne le saisie user
func Scanner(label string) string {
	var tmp string
	fmt.Print(label)
	fmt.Scanln(&tmp)
	return tmp
}