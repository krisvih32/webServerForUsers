package main

import "fmt"

type QueryParamNames struct {
	FirstName string
	LastName  string
	Email     string
	PhoneNo   string
}

type QueryParamNamesInitializer map[string]any

func NewQueryParamNamesInitializer() *QueryParamNamesInitializer {
	return &QueryParamNamesInitializer{}
}

func (main QueryParamNamesInitializer) SetFirstName(firstName string) {
	main["FirstName"] = firstName
}

func (main QueryParamNamesInitializer) SetLastName(lastName string) {
	main["LastName"] = lastName
}

func (main QueryParamNamesInitializer) SetEmail(email string) {
	main["Email"] = email
}

func (main QueryParamNamesInitializer) SetPhoneNo(phoneNo string) {
	main["PhoneNo"] = phoneNo
}

func (main QueryParamNames) GetFirstName() string {
	return main.FirstName
}

func (main QueryParamNames) GetLastName() string {
	return main.LastName
}

func (main QueryParamNames) GetEmail() string {
	return main.Email
}

func (main QueryParamNames) GetPhoneNo() string {
	return main.PhoneNo
}

func (main QueryParamNamesInitializer) NewQueryParamNames() *QueryParamNames {
	queryParamNames := &QueryParamNames{
		fmt.Sprintf("%s", main["FirstName"]),
		fmt.Sprintf("%s", main["LastName"]),
		fmt.Sprintf("%s", main["Email"]),
		fmt.Sprintf("%s", main["PhoneNo"])}
	return queryParamNames
}
