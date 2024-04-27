package service

import (
	"context"
	"errors"
	"roomino/ctl"
	"roomino/dao"
	"roomino/model"
	"roomino/types"
	"sync"

	"gorm.io/gorm"
)

type UserSrv struct {
}

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (s *UserSrv) Register(ctx context.Context, req *types.UserServiceReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	u, err := userDao.FindUserByUserName(req.Username)
	switch err {
	case gorm.ErrRecordNotFound:
		u = &model.Users{
			Username:  req.Username,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			DOB:       req.DOB,
			Gender:    req.Gender,
			Email:     req.Email,
			Phone:     req.Phone,
		}

		if err = u.SetPassword(req.Passwd); err != nil {
			return
		}

		if err = userDao.CreateUser(u); err != nil {
			return
		}

		return ctl.RespSuccess(), nil
	case nil:
		err = errors.New("Userexists")
		return
	default:
		return
	}
}
