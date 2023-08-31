package repository

type Config struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

var (
	config = Config{
		host:     "localhost",
		port:     5432,
		user:     "postgres",
		password: "OK",
		dbname:   "postgres",
	}
)
