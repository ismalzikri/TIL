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

// func sqlQuery() {
// 	db, err := connect()

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	var age = 22

// 	rows, err := db.Query("select id, name, experience from employees where age = $1", age)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer rows.Close()

// 	var result []employee

// 	for rows.Next() {
// 		var each = employee{}
// 		var err = rows.Scan(&each.id, &each.name, &each.experience)

// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return
// 		}

// 		result = append(result, each)
// 	}

// 	if err = rows.Err(); err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	for _, each := range result {
// 		fmt.Println(each.name)
// 	}

// }

// func sqlQueryRow() {
// 	var db, err = connect()

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	var result = employee{}
// 	var id = "1"
// 	err = db.QueryRow("select name, experience from employees where id =$1", id).
// 		Scan(&result.name, &result.experience)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Printf("name: %s\nexperience: %d\n", result.name, result.experience)
// }

func sqlPrepare() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("select name, experience from employees where id=$1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = employee{}
	stmt.QueryRow("1").Scan(&result1.name, &result1.experience)
	fmt.Printf("name: %s\nexperience: %d\n", result1.name, result1.experience)

	var result2 = employee{}
	stmt.QueryRow("2").Scan(&result2.name, &result2.experience)
	fmt.Printf("name: %s\nexperience: %d\n", result2.name, result2.experience)

	var result3 = employee{}
	stmt.QueryRow("3").Scan(&result3.name, &result3.experience)
	fmt.Printf("name: %s\nexperience: %d\n", result3.name, result3.experience)

}

func main() {
	sqlPrepare()
}
