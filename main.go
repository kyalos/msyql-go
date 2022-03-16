package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go SQl")

	db, err := sql.Open("mysql", "kyalos:sealed@/luminous")

	if err != nil {
		fmt.Println("error somewhere")
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("succesfully connected to sql")
}
