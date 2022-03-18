package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Services struct {
	Service_id          int
	Service_name        string
	Service_price       string
	Service_duration    string
	Service_category    string
	Service_description string
	Created_on          string
}

func main() {
	fmt.Println("Go SQl")

	db, err := sql.Open("mysql", "kyalos:sealed@/luminous")

	if err != nil {
		fmt.Println("error somewhere")
		panic(err.Error())
	}

	defer db.Close()

	// insert, err := db.Query("INSERT INTO tester VALUES('ELLIOT')")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	//fmt.Printf("Inserted")

	results, err := db.Query("SELECT * FROM services")

	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var services Services

		err = results.Scan(&services.Service_id, &services.Service_name, &services.Service_price, &services.Service_duration, &services.Service_category, &services.Service_description, &services.Created_on)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(services)
	}
}
