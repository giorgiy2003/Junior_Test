package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type person struct {
	Id        int
	email     string
	phone     string
	firstName string
	lastName  string
}
