package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type employee struct {
	id         string
	name       string
	age        int
	experience int
}

func connect() (*sql.DB, error) {
	connStr := "user=ismaldts dbname=employee sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 22

	rows, err := db.Query("select id, name, experience from employees where age = $1", age)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []employee

	for rows.Next() {
		var each = employee{}
		var err = rows.Scan(&each.id, &each.name, &each.experience)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}

}

func main() {
	sqlQuery()
}
