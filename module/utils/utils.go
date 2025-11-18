package utils

func AddComaOrNot(idx int, lenAttr int) string {
	if idx < lenAttr {
		return ","
	}
	return ""
}