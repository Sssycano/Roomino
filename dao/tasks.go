package dao

import (
	"context"
	"roomino/model"
	"roomino/types"

	"gorm.io/gorm"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewDBClient(ctx)}
}

func (dao *TaskDao) GetUnitsWithPetPolicy(companyName, buildingName, username string) ([]types.UnitInfoResp, error) {
	var units []model.ApartmentUnit
	err := dao.DB.Where("company_name = ? AND building_name = ?", companyName, buildingName).Find(&units).Error
	if err != nil {
		return nil, err
	}

	// 获取用户的宠物信息
	var userPets []model.Pets
	err = dao.DB.Where("username = ?", username).Find(&userPets).Error
	if err != nil {
		return nil, err
	}

	// 获取宠物政策
	var petPolicies []model.PetPolicy // 使用 model 包中的 PetPolicy
	err = dao.DB.Where("company_name = ? AND building_name = ?", companyName, buildingName).Find(&petPolicies).Error
	if err != nil {
		return nil, err
	}

	// 创建宠物政策映射
	petPolicyMap := make(map[string]bool)
	for _, policy := range petPolicies {
		key := policy.PetType + "-" + policy.PetSize
		petPolicyMap[key] = policy.IsAllowed
	}

	var unitInfos []types.UnitInfoResp
	for _, unit := range units {
		isPetAllowed := true // 假定允许宠物

		// 检查每个用户的宠物
		for _, pet := range userPets {
			key := pet.PetType + "-" + pet.PetSize
			if allowed, ok := petPolicyMap[key]; !ok || !allowed { // 如果宠物不被允许
				isPetAllowed = false
				break // 如果任何宠物不被允许，停止检查
			}
		}

		UnitInfoResp := types.UnitInfoResp{
			UnitRentID:             unit.UnitRentID,
			MonthlyRent:            unit.MonthlyRent,
			SquareFootage:          unit.SquareFootage,
			AvailableDateForMoveIn: unit.AvailableDateForMoveIn,
			IsPetAllowed:           isPetAllowed,
		}

		unitInfos = append(unitInfos, UnitInfoResp) // 添加单元信息
	}

	return unitInfos, nil
}
