package repositories

import "github.com/jinzhu/gorm"

type BaseRepository interface {
	CreateUser(db *gorm.DB, entity interface{}) error
}
