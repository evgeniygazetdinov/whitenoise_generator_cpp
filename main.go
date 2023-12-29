package main

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func dbConnect() {

	db, err := sql.Open("mysql", "root:my-secret-pw@/tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("CONNECT TO DB")

	insert, err := db.Query("INSERT into  users (name, age) values('колек', 24)")
	if err != nil {
		panic(err)
	}
	defer insert.Close()

}

func main() {

	dbConnect()

}
