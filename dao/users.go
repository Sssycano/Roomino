package dao

import (
	"context"

	"roomino/model"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) FindUserByUserName(userName string) (user *model.Users, err error) {
	err = dao.DB.Model(&model.Users{}).Where("Username=?", userName).
		First(&user).Error

	return
}

func (dao *UserDao) CreateUser(user *model.Users) (err error) {
	err = dao.DB.Model(&model.Users{}).Create(user).Error
	return
}
