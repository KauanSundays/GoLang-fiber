package storage

import(
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host		string
	Port		string
	Password	string
	User		string
	DBName		string
	SSLMode		string
}

func NewConnection(config *Config)(*gorm.DB, error) {
	dsn : fmt.Sprintf(
		"host=%s  port=%s user=%s password=%s dbname=%s sslmode=%s", 
	)
}