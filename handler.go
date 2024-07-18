package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	db              *sql.DB
	connectionData  Credentials
	queryParamNames QueryParamNames
	errorMessage    string
}

func NewHandler(
	db *sql.DB,
	connectionData Credentials,
	queryParamNames QueryParamNames,
) *Handler {
	return &Handler{
		db:              db,
		connectionData:  connectionData,
		queryParamNames: queryParamNames,
	}
}

func (handler Handler) POSTHandler(context echo.Context) error {
	firstName, lastName := context.QueryParams().Get(handler.queryParamNames.FirstName), context.QueryParams().Get(handler.queryParamNames.LastName)
	if firstName == "" {
		log.Printf("err is firstName required and context is %v and connectionString is %s", context, handler.connectionData.GetConnectionString())
		return context.String(http.StatusBadRequest, "first name required")
	}
	if lastName == "" {
		return context.String(http.StatusBadRequest, "last name required")
	}
	_, err := handler.db.Exec("DELETE FROM addressBook WHERE firstName =? AND lastName =?", firstName, lastName)
	if err != nil {
		log.Printf("err is %v and context is %v and connectionString is %s", err, context, handler.connectionData.GetConnectionString())
		return context.String(http.StatusInternalServerError, "")
	}
	return context.String(http.StatusOK, "")
}

func (handler Handler) GETHandler(context echo.Context) error {
	var firstName, lastName string = context.QueryParams().Get("firstName"), context.QueryParams().Get("lastName")
	if firstName != "" && lastName != "" {
		db, err := sql.Open("mysql", handler.connectionData.GetConnectionString())
		if err != nil {
			log.Printf("err is %v and context is %v and connectionString is %s", err, context, handler.connectionData.GetConnectionString())
			return context.String(http.StatusInternalServerError, "")
		}
		defer db.Close()
		row, err := db.Query("SELECT * FROM addressBook WHERE firstName =? AND lastName =?", firstName, lastName)
		if err != nil {
			log.Printf("err is %v and context is %v and connectionString is %s", err, context, handler.connectionData.GetConnectionString())
			return context.String(http.StatusInternalServerError, "")
		}
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
		log.Printf("err is %v and context is %v and connectionString is %s", err, context, handler.connectionData.GetConnectionString())
		return context.String(http.StatusInternalServerError, "")
	}
	return context.JSON(http.StatusOK, rows)
}

func (handler Handler) Error() string {
	return handler.errorMessage
}

func adaptHandler(
	db *sql.DB,
	connectionData Credentials,
	QueryParamNames QueryParamNames) echo.HandlerFunc {
	return func(c echo.Context) error {
		return NewHandler(db, connectionData, QueryParamNames)
	}
}
