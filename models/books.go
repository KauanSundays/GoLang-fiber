package models //model

import "gorm.io/gorm"

type Books  struct {
	ID 			 uint  `gorm: "primary key;autoIncrement" json:"id"` //Other options in gorm doc
	Author 		*string    `json:"author"`
	Title 		*string    `json:"title"`
	Publisher	*string    `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error{ // migrate function

}