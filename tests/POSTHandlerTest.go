package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"sqlmock"
)

// POSTHandler should return 200 OK when all required query parameters are provided and database operations succeed
func TestPOSTHandlerSuccessCorrectUsage(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	connectionData := Credentials{connectionString: "test_connection_string"}
	queryParamNames := QueryParamNames{
		FirstName: "firstName",
		LastName:  "lastName",
		Email:     "email",
		PhoneNo:   "phoneNo",
	}

	handler := NewHandler(db, connectionData, queryParamNames)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/addressBookWebService?firstName=John&lastName=Doe&email=john.doe@example.com&phoneNo=1234567890", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mock.ExpectPing().WillReturnError(nil)
	mock.ExpectExec("DELETE FROM addressBook WHERE firstName =? AND lastName =?").
		WithArgs("John", "Doe").
		WillReturnResult(sqlmock.NewResult(1, 1))

	if assert.NoError(t, handler.POSTHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}
}

// POSTHandler should return 400 Bad Request when firstName query parameter is missing
func TestPOSTHandlerMissingFirstName(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	connectionData := Credentials{connectionString: "test_connection_string"}
	queryParamNames := QueryParamNames{
		FirstName: "firstName",
		LastName:  "lastName",
		Email:     "email",
		PhoneNo:   "phoneNo",
	}

	handler := NewHandler(db, connectionData, queryParamNames)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/addressBookWebService?lastName=Doe&email=john.doe@example.com&phoneNo=1234567890", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.POSTHandler(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, "first name required", rec.Body.String())
	}
}
