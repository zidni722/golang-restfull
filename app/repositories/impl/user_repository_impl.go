package impl

import (
	"github.com/jinzhu/gorm"
	"github.com/zidni722/golang-restfull/app/models"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) CreateUser(db *gorm.DB, entity interface{}) error {
	return db.Create(entity.(*models.User)).Error
}

func (r *UserRepositoryImpl) FindAll(db *gorm.DB, entities interface{}) error {
	return db.Find(entities.(*[]models.User)).Error
}

func (r *UserRepositoryImpl) FindById(db *gorm.DB, entity interface{}, id int) error {
	return db.First(entity.(*models.User), id).Error
}

func (r *UserRepositoryImpl) Create(db *gorm.DB, entity interface{}) error {
	return db.Create(entity.(*models.User)).Error
}

func (r *UserRepositoryImpl) NewRecord(db *gorm.DB, entity interface{}) bool {
	return db.NewRecord(entity.(models.User))
}

func (r *UserRepositoryImpl) Update(db *gorm.DB, entity interface{}) error {
	return db.Save(entity.(*models.User)).Error
}

func (r *UserRepositoryImpl) Delete(db *gorm.DB, entity interface{}) error {
	return db.Delete(entity.(*models.User)).Error
}
