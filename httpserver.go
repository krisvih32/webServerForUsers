package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/tabwriter"

	"github.com/labstack/echo/v4"
)

const (
	FirstName = "firstName"
	LastName  = "lastName"
)

// print sql table to writer in form of `column: values`
func printTable(writer io.Writer, table *sql.Rows) error {
	columns, err := table.Columns()
	if err != nil {
		return err
	}
	for table.Next() {
		values := make([]any, len(columns))
		scanArguments := make([]any, len(values))
		// copy values from values to scanArgs
		for i := range values {
			scanArguments[i] = &values[i]
		}
		if err := table.Scan(scanArguments...); err != nil {
			return err
		}
		tabwriter := tabwriter.NewWriter(writer, 0, 4, 2, byte('\t'), 0)
		for _, column := range columns {
			fmt.Fprintf(tabwriter, "\t"+column)
		}
		for row := range columns {
			fmt.Fprintf(tabwriter, "\t%v", values[row])
		}
	}
	return nil
}

func POSTHandler(context echo.Context) error {
	firstName, lastName := context.QueryParams().Get(FirstName), context.QueryParams().Get(LastName)
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
		log.Printf("%v", err)
		return context.String(http.StatusInternalServerError, "")
	}
	defer db.Close()
	// If the first name and last name already exist in the database, delete them
	_, err = db.Exec("DELETE FROM addressBook WHERE firstName =? AND lastName =?", firstName, lastName)
	if err != nil {
		log.Printf("%v", err)
		return context.String(http.StatusInternalServerError, "")
	}
	return context.String(http.StatusOK, "")
}

func GETHandler(context echo.Context) error {
	var firstName, lastName string = context.QueryParams().Get("firstName"), context.QueryParams().Get("lastName")
	//check first name and last name are inputted, if inputted, return table where first name and last name match
	// if none match, return empty table
	if firstName != "" && lastName != "" {
		// Open a connection to the database
		db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/addressBookWebService")
		if err != nil {
			log.Printf("%v", err)
			return context.String(http.StatusInternalServerError, "")
		}
		defer db.Close()
		row, err := db.Query("SELECT * FROM addressBook WHERE firstName =? AND lastName =?", firstName, lastName)
		if err != nil {
			log.Printf("%v", err)
			return context.String(http.StatusInternalServerError, "")
		}
		var writer io.Writer
		err = printTable(writer, row)
		if err != nil {
			fmt.Printf("%v", err)
			return context.String(http.StatusInternalServerError, "")
		}
		output := fmt.Sprintf("%#v", writer)
		return context.String(http.StatusOK, output)
	}
	// Open a connection to the database
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/addressBookWebService")
	if err != nil {
		fmt.Printf("%v", err)
		return context.String(http.StatusInternalServerError, "")
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM addressBook")
	if err != nil {
		log.Printf("%v", err)
		return context.String(http.StatusInternalServerError, "")
	}
	var writer io.Writer
	err = printTable(writer, rows)
	if err != nil {
		fmt.Printf("%v", err)
		return context.String(http.StatusInternalServerError, "")
	}
	output := fmt.Sprintf("%#v", writer)
	return context.String(http.StatusOK, output)
}
