package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func POSTHandler(context echo.Context) error {
	var firstName, lastName string = context.QueryParams().Get("firstName"), context.QueryParams().Get("lastName")
	//check first name and last name are inputted
	if firstName == "" {
		return context.String(http.StatusBadRequest, "first name required")
	}
	if lastName == "" {
		return context.String(http.StatusBadRequest, "last name required")
	}
	// Open a connection to the database
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/addressBookWebService")
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	defer db.Close()
	// If the first name and last name already exist in the database, delete them
	_, err = db.Exec("DELETE FROM addressBook WHERE firstName =? AND lastName =?", firstName, lastName)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return context.String(http.StatusOK, "")
}
