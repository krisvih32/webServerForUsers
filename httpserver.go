package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

var jsonObj []any

var queryParamNames QueryParamNames

type Handler struct {
	db             *sql.DB
	connectionData Credentials
}

type QueryParamNames struct {
	FirstName string
	LastName  string
	Email     string
	PhoneNo   string
}

type PrintingStruct struct {
	MinWidth int
	TabWidth int
	Padding  int
	PadChar  byte
	Flags    uint
}

func newPrinting() PrintingStruct {
	return PrintingStruct{
		MinWidth: 0,
		TabWidth: 4,
		Padding:  2,
		PadChar:  byte('\t'),
		Flags:    0,
	}
}

var jsonObj []any

var queryParamNames QueryParamNames

type Handler struct {
	db             *sql.DB
	connectionData Credentials
}

type QueryParamNames struct {
	FirstName string
	LastName  string
	Email     string
	PhoneNo   string
}

type PrintingStruct struct {
	MinWidth int
	TabWidth int
	Padding  int
	PadChar  byte
	Flags    uint
}

func newPrinting() PrintingStruct {
	return PrintingStruct{
		MinWidth: 0,
		TabWidth: 4,
		Padding:  2,
		PadChar:  byte('\t'),
		Flags:    0,
	}
}

const (
	BaseStatusInternalServerErrorString = `If you see this, something is wrong with the server.
	Please try again.
	If error proceeds, contact the developer.
	This info might be useful: error code: `
)

func makeBaseStatusInternalServerErrorResponse(code int) string {
	return fmt.Sprintf("%s%d", BaseStatusInternalServerErrorString, code)
}

func adaptHandler(handler func(handler Handler, context echo.Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handler(Handler{}, c)
	}
}

func (handler Handler) POSTHandler(context echo.Context) error {
	firstName, lastName := context.QueryParams().Get(queryParamNames.FirstName), context.QueryParams().Get(queryParamNames.LastName)
}

func makeBaseStatusInternalServerErrorResponse(code int) string {
	return fmt.Sprintf("%s%d", BaseStatusInternalServerErrorString, code)
}

func adaptHandler(handler func(handler Handler, context echo.Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handler(Handler{}, c)
	}
}

func (handler Handler) POSTHandler(context echo.Context) error {
	firstName, lastName := context.QueryParams().Get(queryParamNames.FirstName), context.QueryParams().Get(queryParamNames.LastName)
	if firstName == "" {
		return context.String(http.StatusBadRequest, "first name required")
	}
	if lastName == "" {
		return context.String(http.StatusBadRequest, "last name required")
	}
	_, err := handler.db.Exec("DELETE FROM addressBook WHERE firstName =? AND lastName =?", firstName, lastName)
	if err != nil {
		log.Printf("%v", err)
		return context.String(http.StatusInternalServerError, "")
	}
	return context.String(http.StatusOK, "")
}

func (handler Handler) GETHandler(context echo.Context) error {
	var firstName, lastName string = context.QueryParams().Get("firstName"), context.QueryParams().Get("lastName")
	if firstName != "" && lastName != "" {
		db, err := sql.Open("mysql", handler.connectionData.GetConnectionString())
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
		defer row.Close()
		return context.JSON(http.StatusOK, row)
		defer row.Close()
		return context.JSON(http.StatusOK, row)
	}
	db, err := sql.Open("mysql", handler.connectionData.GetConnectionString())
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
	return context.JSON(http.StatusOK, rows)
}

type exitCodes map[string]int

func main() {
	connectionStringInitializer := NewCredentialsInitializer()
	connectionStringInitializer.SetDatabase("mysql")
	connectionStringInitializer.SetUsername("username")
	connectionStringInitializer.SetPassword("pass")
	connectionStringInitializer.SetHostname("192.168.68.103")
	connectionStringInitializer.SetPort(8080)
	connectionStringInitializer.SetUsername("username")
	connectionStringInitializer.SetEmail("email")
	connectionStringInitializer.SetFirstName("firstName")
	connectionStringInitializer.SetLastName("lastName")
	connectionString := connectionStringInitializer.NewCredentials()
	// exit scope of connectionStringInitializer because it is not used
	{ // create a new Echo instance
		e := echo.New()
		defer e.Close()
		e.POST("/addressBookWebService", adaptHandler(Handler.POSTHandler))
		e.GET("/addressBookWebService", adaptHandler(Handler.GETHandler))
		slog.Debug(fmt.Sprintf("%v", connectionString))
		e.Logger.Fatalf("%v", e.Start(fmt.Sprintf("%s:%s", connectionString.hostname, connectionString.port)))
	}
}
