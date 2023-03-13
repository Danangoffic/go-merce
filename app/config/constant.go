package config

// Setter
const (
	// Set the Database Username
	UserDB string = "root"
	// Set the Database Password
	PassDB string = ""
	// Set The Database Connection Type
	ConnectionDB string = "tcp"
	// Set the Databse Host
	HostDB string = "127.0.0.1"
	// Set the Database Port
	PortDB string = "3306"
	// Set the Database Name
	DBName string = "ecommerce_go"
	// Set the Database optional Extra
	DBOptionExtra string = "charset=utf8mb4&parseTime=True&loc=Local"
)

// Getter
const (
	// Get the DSN Database Configuration, it will return as a String. Just Change The Setter!
	DSN string = UserDB + ":" + PassDB + "@" + PortDB + "(" + HostDB + ":" + PortDB + ")/" + DBName + "?" + DBOptionExtra
)
