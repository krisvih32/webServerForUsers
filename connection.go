package main

import "fmt"

type Credentials struct {
	firstName        string
	lastName         string
	email            string
	phoneNo          string
	username         string
	password         string
	SQLhostname      string
	serverHostname   string
	port             string
	database         string
	connectionString string
}
type CredentialsInitializer map[string]any

func NewCredentialsInitializer() *CredentialsInitializer {
	return &CredentialsInitializer{}
}
func (credInit CredentialsInitializer) NewCredentials() *Credentials {
	credInit["connectionString"] = credInit.GetConnectionString()
	credentials := &Credentials{
		fmt.Sprintf("%v", credInit["firstName"]),
		fmt.Sprintf("%v", credInit["lastName"]),
		fmt.Sprintf("%v", credInit["email"]),
		fmt.Sprintf("%v", credInit["phoneNo"]),
		fmt.Sprintf("%v", credInit["username"]),
		fmt.Sprintf("%v", credInit["password"]),
		fmt.Sprintf("%v", credInit["SQLhostname"]),
		fmt.Sprintf("%v", credInit["serverHostname"]),
		fmt.Sprintf("%v", credInit["port"]),
		fmt.Sprintf("%v", credInit["database"]),
		fmt.Sprintf("%v", credInit["connectionString"])}
	return credentials
}


func (main CredentialsInitializer) SetQueryParamNames(queryParamNames QueryParamNames){
	main["queryParamNames"] = queryParamNames
}

func (main CredentialsInitializer) SetFirstName(firstName string) {
	main["firstName"] = firstName
}

func (credInit CredentialsInitializer) SetLastName(lastName string) {
	credInit["lastName"] = lastName
}

func (credInit CredentialsInitializer) SetEmail(email string) {
	credInit["email"] = email
}

func (credInit CredentialsInitializer) SetPhoneNo(phoneNo string) {
	credInit["phoneNo"] = phoneNo
}

func (credInit CredentialsInitializer) SetUsername(username string) {
	credInit["username"] = username
}
func (credInit CredentialsInitializer) SetPassword(password string) {
	credInit["password"] = password
}

func (credInit CredentialsInitializer) SetSQLHostname(hostname string) {
	credInit["SQLhostname"] = hostname
}
func (credInit CredentialsInitializer) SetServerHostname(hostname string) {
	credInit["serverHostname"] = hostname
}
func (credInit CredentialsInitializer) SetPort(port int) {
	credInit["port"] = port
}

func (credInit CredentialsInitializer) SetDatabase(database string) {
	credInit["database"] = database
}

func (cred Credentials) GetFirstName() string {
	return cred.firstName
}

func (cred Credentials) GetLastName() string {
	return cred.lastName
}

func (cred Credentials) GetEmail() string {
	return cred.email
}

func (cred Credentials) GetPhoneNo() string {
	return cred.phoneNo
}

func (cred Credentials) GetUsername() string {
	return cred.username
}

func (cred Credentials) GetPassword() string {
	return cred.password
}

func (cred Credentials) GetSQLHostname() string {
	return cred.SQLhostname
}

func (cred Credentials) GetServerHostname() string {
	return cred.serverHostname
}

func (cred Credentials) GetPort() string {
	return cred.port
}

func (cred Credentials) GetDatabase() string {
	return cred.database
}

func (cred CredentialsInitializer) GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/addressBookWebService", cred["username"], cred["password"], cred["hostname"], cred["port"])
}
func (cred Credentials) GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/addressBookWebService", cred.username, cred.password, cred.serverHostname, cred.port)
}
