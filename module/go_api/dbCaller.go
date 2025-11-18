package goapi

func DbCallerPG() string {
	return `
db, err := sql.Open("postgres", database_url)
if err != nil {
	log.Fatal(err)
}
defer db.Close()`
}

func DBCallerHandler(sgbd string) string {
	if sgbd =="postgresql" {
		return DbCallerPG()
	}
	return ""
}