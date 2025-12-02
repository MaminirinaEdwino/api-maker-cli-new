package apimaker

import "fmt"

func GetSGBD() string {
	var sgbd string
	fmt.Println("Choose your SGBD : ")
	fmt.Println("1. postgresql")
	fmt.Println("2. mysql")
	fmt.Scan(&sgbd)
	switch sgbd {
	case "postgresql":
		return sgbd
	case "mysql":
		return sgbd
	}
	return GetSGBD()
}