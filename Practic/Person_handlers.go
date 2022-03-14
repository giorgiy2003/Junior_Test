package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)


func home_page(w http.ResponseWriter, r *http.Request) {
	database, _ := sql.Open("sqlite3", "./persons.db") //Для открытия соединения с базой данных используем функцию sql.Open()
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS persons (Id INTEGER PRIMARY KEY,email TEXT,phone TEXT,firstName TEXT,lastName TEXT)")
	statement.Exec()

	//statement, _ = database.Prepare("INSERT INTO persons (email,phone,firstName,lastName) VALUES (?,?,?,?)")
	//statement.Exec("Masha@email", "+79545454767", "masha", "veselova")

	rows, _ := database.Query("SELECT Id,email,phone,firstName,lastName FROM persons")

	persons := []person{}

	for rows.Next() {
		p := person{}
		rows.Scan(&p.Id, &p.email, &p.phone, &p.firstName, &p.lastName)
		persons = append(persons, p)
	}

	for _, p := range persons {
		fmt.Fprintln(w, p.formatStyle())
	}
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8090", nil)
}

func main() {
	handleRequest()
}

