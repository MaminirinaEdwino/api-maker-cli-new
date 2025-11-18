package postgres

func DatabaseTypeConverter(attr_type string) string {
	switch attr_type {
	case "int":
		return "INTEGER"
	case "string":
		return "VARCHAR(255)"
	}
	return ""
}