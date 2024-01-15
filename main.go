package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const PORT = ":8080"

type ViewData struct {
	Title   string
	Message string
}

var database *sql.DB

type Products struct {
	Id      int
	Model   string
	Company string
	Price   int
}

func handleFunc() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		rows, err := database.Query("select * from products")
		if err != nil {
			fmt.Println("error")
		}
		defer rows.Close()
		products := []Products{}
		for rows.Next() {
			p := Products{}
			err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
			if err != nil {
				fmt.Println(err)
				continue
			}
			products = append(products, p)
		}

		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, products)
	})
	fmt.Printf("Running on %s \n", PORT)
	http.ListenAndServe(PORT, nil)

}
func main() {
	db, err := sql.Open("mysql", "docker:password@tcp(0.0.0.0:3306)/golang")

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()
	handleFunc()
}
