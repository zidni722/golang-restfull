package repositories

import "github.com/jinzhu/gorm"

type BaseRepository interface {
	FindAll(db *gorm.DB, entities interface{}) error
	FindById(db *gorm.DB, entity interface{}, id int) error
	Create(db *gorm.DB, entity interface{}) error
	NewRecord(db *gorm.DB, entity interface{}) bool
	Update(db *gorm.DB, entity interface{}) error
	Delete(db *gorm.DB, entity interface{}) error
}
