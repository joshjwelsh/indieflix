package config

import "fmt"

type DBConfig struct {
	user     string
	password string
	host     string
	dbName   string
}

func (c DBConfig) ConnectionString() string {
	port := "5432"
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.host, port, c.user, c.password, c.dbName)
}

func NewDBConfig(user, password, host, dbName string) DBConfig {
	return DBConfig{
		user:     user,
		password: password,
		host:     host,
		dbName:   dbName,
	}
}
