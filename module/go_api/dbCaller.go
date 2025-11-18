package goapi

func DbCallerPG() string {
	return `
db, err := sql.Open("postgres", database_url)
if err != nil {
	log.Fatal(err)
}
defer db.Close()`
}