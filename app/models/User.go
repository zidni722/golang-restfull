package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UUID    string `gorm:"column:uuid"`
	Name    string
	Address string
	Gender  string
}
