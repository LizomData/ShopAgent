package model

type Config struct {
	DatabaseConfig `ini:"database"`
}

type DatabaseConfig struct {
	Db       string `ini:"db"`
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Name     string `ini:"name"`
}
