package service

import (
	"context"
	"errors"
	"roomino/ctl"
	"roomino/dao"
	"roomino/types"
	"sync"
)

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

type TaskSrv struct {
}

func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

func (s *TaskSrv) GetAvailableUnitsWithPetPolicy(ctx context.Context, req *types.UnitInforeq) (resp interface{}, err error) {
	// 创建 TaskDao 实例
	taskDao := dao.NewTaskDao(ctx)
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return
	}
	// 从 DAO 获取出租单位和宠物政策
	units, err := taskDao.GetUnitsWithPetPolicy(req.CompanyName, req.BuildingName, u.UserName)
	if err != nil {
		return nil, errors.New("failed to retrieve units") // 错误处理
	}

	// 转换结果为 UnitInfoResp 类型
	var unitResp []types.UnitInfoResp
	for _, unit := range units {
		unitResp = append(unitResp, types.UnitInfoResp{
			UnitRentID:             unit.UnitRentID,
			MonthlyRent:            unit.MonthlyRent,
			SquareFootage:          unit.SquareFootage,
			AvailableDateForMoveIn: unit.AvailableDateForMoveIn,
			IsPetAllowed:           unit.IsPetAllowed,
		})
	}

	// 返回成功响应
	return ctl.RespSuccessWithData(unitResp), nil
}
