package go46_database

var connection string

func init() { /* <- init function it's automatically trigger without call the function */
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
