package config

const (
	UserDB        string = "root"
	PassDB        string = ""
	ConnectionDB  string = "tcp"
	HostDB        string = "127.0.0.1"
	PortDB        string = "3306"
	DBName        string = "ecommerce_go"
	DBOptionExtra string = "charset=utf8mb4&parseTime=True&loc=Local"
)

const (
	DSN string = UserDB + ":" + PassDB + "@" + PortDB + "(" + HostDB + ":" + PortDB + ")/" + DBName + "?" + DBOptionExtra
)
