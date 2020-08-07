package config

// DBConfig holds database connection's credentials.
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
}
