package main

type Config struct {
	Dsn string
}

func LoadConfig() *Config {
	return &Config{
		Dsn: "postgresql://postgres:postgres1234@localhost/decode?sslmode=disable",
	}
}
