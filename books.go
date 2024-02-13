package models

import "gorm.io/gorm"

type Books  struct {
	ID
	Author
	Title
	Publisher
}