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

func (dao *UserDao) CreateUser(user *model.Users) (err error) {
	err = dao.DB.Model(&model.Users{}).Create(user).Error

	return
}
