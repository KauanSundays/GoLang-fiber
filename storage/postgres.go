package storage

import(
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host
	Port
	Password
	User
	DBName
	SSLMode
}