package util

type DbConfig struct {
	DbName       string
	Host         string
	Password     string
	Port         string
	User         string
	Schema       string
	LogFile      string
	MaxOpenConns int
	MaxIdleConns int
}
