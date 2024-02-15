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
	SSLMode	
}

func NewConnection(config *Config)(*gorm.DB, error) {
	
}