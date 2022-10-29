package db

type DatabaseList struct {
	EMoneySvc Database
}

type Database struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}
