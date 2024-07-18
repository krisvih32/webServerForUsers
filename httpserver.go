package main

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/labstack/echo/v4"
)

type PrintingStruct struct {
	MinWidth int
	TabWidth int
	Padding  int
	PadChar  byte
	Flags    uint
}

const (
	BaseStatusInternalServerErrorString = `If you see this, something is wrong with the server.
	Please try again.
	If error proceeds, contact the developer.
	This info might be useful: error code: `
)

// Initializes a queryParamNamesInitializer and converts it into queryParamNames
// Initializes a connectionStringInitializer and converts it into a connectionString
// Creates a mysql sql.DB object
// If it fails to create db object it returns 1
// Creates a handler object using db, queryParamNames and connectionString together
// Creates an echo object and defers it
// Setting up POSTHandler and GETHandler to run when path /addressBookWebService is opened with POST or GET respectively
// Starts web service
// Return Values:
// 		0: success
// 		1: db couldn't be initialized

func main() {
	queryParamNamesInitializer := NewQueryParamNamesInitializer()
	queryParamNamesInitializer.SetFirstName("firstName")
	queryParamNamesInitializer.SetLastName("lastName")
	queryParamNamesInitializer.SetEmail("email")
	queryParamNamesInitializer.SetPhoneNo("phoneNo")
	queryParamNames := queryParamNamesInitializer.NewQueryParamNames()
	connectionStringInitializer := NewCredentialsInitializer()
	connectionStringInitializer.SetDatabase("mysql")
	connectionStringInitializer.SetUsername("username")
	connectionStringInitializer.SetPassword("pass")
	connectionStringInitializer.SetServerHostname("localhost")
	connectionStringInitializer.SetPort(8080)
	connectionStringInitializer.SetUsername("username")
	connectionStringInitializer.SetEmail("email")
	connectionStringInitializer.SetFirstName("firstName")
	connectionStringInitializer.SetLastName("lastName")
	connectionStringInitializer.SetPhoneNo("phoneNo")
	connectionStringInitializer.SetSQLHostname("DESKTOP-D4U24AV")
	connectionStringInitializer.SetQueryParamNames(*queryParamNames)
	connectionString := connectionStringInitializer.NewCredentials()
	db, err := sql.Open("mysql", connectionString.GetConnectionString())
	if (err != nil){
		os.Exit(1)
	}
	handler:=NewHandler(db, *connectionString, *queryParamNames)
	e := echo.New()
	defer e.Close()
	e.POST("/addressBookWebService", adaptHandler(db, handler.connectionData, *queryParamNames))
	e.GET("/addressBookWebService", adaptHandler(db, handler.connectionData, *queryParamNames))
	address := fmt.Sprintf("%s:%s", connectionString.GetServerHostname(), connectionString.GetPort())
	e.Logger.Fatalf("%v", e.Start(address))
}

