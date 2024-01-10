package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json: "name"`
	Age  uint16 `json: "age"`
}

func insertUserIntoDb(db *sql.DB) {
	insert, err := db.Query("INSERT into  users (name, age) values('kolek', 24)")
	defer insert.Close()
	if err != nil {
		panic(err)
	}
}

func printAllUsersInsideDb(db *sql.DB) {

	res, err := db.Query("select name, age from `users`")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("User name: %s age: %d", user.Name, user.Age))
	}
}

func dbConnect() *sql.DB {

	db, err := sql.Open("mysql", "docker:password@tcp(0.0.0.0:3306)/golang")
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	return db
}

func main() {

	db := dbConnect()
	insertUserIntoDb(db)
	printAllUsersInsideDb(db)
}
