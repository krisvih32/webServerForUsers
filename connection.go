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
func (self CredentialsInitializer) NewCredentials() *Credentials {
	self["connectionString"] = self.GetConnectionString()
	credentials := &Credentials{
		fmt.Sprintf("%v", self["database"]),
		fmt.Sprintf("%v", self["username"]),
		fmt.Sprintf("%v", self["email"]),
		fmt.Sprintf("%v", self["phoneNo"]),
		fmt.Sprintf("%v", self["firstName"]),
		fmt.Sprintf("%v", self["lastName"]),
		fmt.Sprintf("%v", self["password"]),
		fmt.Sprintf("%v", self["SQLhostname"]),
		fmt.Sprintf("%v", self["serverHostname"]),
		fmt.Sprintf("%v", self["port"]),
		""}
	return credentials
}

func (self CredentialsInitializer) SetFirstName(firstName string) {
	self["firstName"] = firstName
}

func (self CredentialsInitializer) SetLastName(lastName string) {
	self["lastName"] = lastName
}

func (self CredentialsInitializer) SetEmail(email string) {
	self["email"] = email
}

func (self CredentialsInitializer) SetPhoneNo(phoneNo string) {
	self["phoneNo"] = phoneNo
}

func (self CredentialsInitializer) SetUsername(username string) {
	self["username"] = username
}
func (self CredentialsInitializer) SetPassword(password string) {
	self["password"] = password
}

func (self CredentialsInitializer) SetSQLHostname(hostname string) {
	self["SQLhostname"] = hostname
}
func (self CredentialsInitializer) SetServerHostname(hostname string) {
	self["serverHostname"] = hostname
}
func (self CredentialsInitializer) SetPort(port int) {
	self["port"] = fmt.Sprintf("%d", port)
}

func (self CredentialsInitializer) SetDatabase(database string) {
	self["database"] = database
}

func (self Credentials) GetFirstName() string {
	return self.firstName
}

func (self Credentials) GetLastName() string {
	return self.lastName
}

func (self Credentials) GetEmail() string {
	return self.email
}

func (self Credentials) GetPhoneNo() string {
	return self.phoneNo
}

func (self Credentials) GetUsername() string {
	return self.username
}

func (self Credentials) GetPassword() string {
	return self.password
}

func (self Credentials) GetSQLHostname() string {
	return self.SQLhostname
}

func (self Credentials) GetServerHostname() string {
	return self.serverHostname
}

func (self Credentials) GetPort() string {
	return self.port
}

func (self Credentials) GetDatabase() string {
	return self.database
}

func (self CredentialsInitializer) GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/addressBookWebService", self["username"], self["password"], self["hostname"], self["port"])
}
func (self Credentials) GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/addressBookWebService", self.username, self.password, self.serverHostname, self.port)
}
