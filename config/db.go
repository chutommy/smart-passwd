package config

// DBConfig holds database connection's credentials.
type DBConfig struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	DBname   string `yaml:"Database_Name"`
}
