package cmd

import "os"

type Env struct {
	Environment string
	Dsn         string
}

func GetEnvironmentVariables() Env {
	return Env{
		Environment: os.Getenv("APP_ENV"),
		Dsn:         os.Getenv("APP_DB_DNS"),
	}
}
