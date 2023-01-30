package database

var connection string

func init() { // function yg diekseskusi pertama kali saat package di load (like constructor)
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
