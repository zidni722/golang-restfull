package _interface

import (
	"github.com/jinzhu/gorm"
	"github.com/zidni722/golang-restfull/app/repositories"
)

type IUserRepository interface {
	repositories.BaseRepository
	CreateUser(db *gorm.DB, entity interface{}) error
}
