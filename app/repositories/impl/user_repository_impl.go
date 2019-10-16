package impl

import (
	"github.com/jinzhu/gorm"
	"github.com/zidni722/golang-restfull/app/models"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) CreateUser(db *gorm.DB, entity interface{}) error {
	return db.Create(entity.(*models.User)).Error
}
