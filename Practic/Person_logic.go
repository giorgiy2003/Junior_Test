package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)



func (p person) formatStyle() string {
	return fmt.Sprintf("%d. Person name is %s. His lastname is %s. His phone is %s. Also email: %s", p.Id, p.firstName, p.lastName, p.phone, p.email)
}

