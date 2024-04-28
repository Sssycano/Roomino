package dao

import (
	"context"
	"errors"
	"regexp"
	"roomino/model"
	"roomino/types"
	"strings"

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
	var userPets []model.Pets
	err = dao.DB.Where("username = ?", username).Find(&userPets).Error
	if err != nil {
		return nil, err
	}

	var petPolicies []model.PetPolicy
	err = dao.DB.Where("company_name = ? AND building_name = ?", companyName, buildingName).Find(&petPolicies).Error
	if err != nil {
		return nil, err
	}

	petPolicyMap := make(map[string]bool)
	for _, policy := range petPolicies {
		key := policy.PetType + "-" + policy.PetSize
		petPolicyMap[key] = policy.IsAllowed
	}

	var unitInfos []types.UnitInfoResp
	for _, unit := range units {
		isPetAllowed := true

		for _, pet := range userPets {
			key := pet.PetType + "-" + pet.PetSize
			if allowed, ok := petPolicyMap[key]; !ok || !allowed {
				isPetAllowed = false
				break
			}
		}

		UnitInfoResp := types.UnitInfoResp{
			UnitRentID:             unit.UnitRentID,
			MonthlyRent:            unit.MonthlyRent,
			SquareFootage:          unit.SquareFootage,
			AvailableDateForMoveIn: unit.AvailableDateForMoveIn,
			IsPetAllowed:           isPetAllowed,
		}

		unitInfos = append(unitInfos, UnitInfoResp)
	}

	return unitInfos, nil
}

func (dao *TaskDao) UpdatePet(req *types.UpdatePets, username string) error {
	var pet model.Pets
	err := dao.DB.Where("pet_name = ? AND pet_type = ? AND username = ?", req.CurrentPetName, req.CurrentPetType, username).First(&pet).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return err
	}

	updateData := map[string]interface{}{
		"pet_name": req.NewPetName,
		"pet_type": req.NewPetType,
		"pet_size": req.NewPetSize,
	}

	if err := dao.DB.Model(&pet).Updates(updateData).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") || strings.Contains(err.Error(), "unique constraint") {
			return errors.New("DUPLICATE_KEY: pet already exists")
		}
		return err
	}

	return nil
}
func (dao *TaskDao) GetPet(username string) ([]model.Pets, error) {
	var pets []model.Pets
	if err := dao.DB.Where("username = ?", username).Find(&pets).Error; err != nil {
		return nil, err
	}

	return pets, nil
}

func (dao *TaskDao) CreatePet(req *types.GetPets, username string) error {
	newPet := model.Pets{
		PetName:  req.CurrentPetName,
		PetType:  req.CurrentPetType,
		PetSize:  req.CurrentPetSize,
		Username: username,
	}
	if err := dao.DB.Create(&newPet).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") || strings.Contains(err.Error(), "unique constraint") {
			return errors.New("DUPLICATE_KEY: pet already exists")
		}
		return err
	}

	return nil
}

func (dao *TaskDao) GetInterests(unitRentID int) ([]model.Interests, error) {
	var interests []model.Interests
	if err := dao.DB.Where("unit_rent_id = ?", unitRentID).Find(&interests).Error; err != nil {
		return nil, err
	}
	return interests, nil
}

func (dao *TaskDao) CreateInterests(req *types.PostInterestReq, username string) error {
	newInterest := model.Interests{
		Username:    username,
		UnitRentID:  req.UnitRentID,
		RoommateCnt: req.RoommateCnt,
		MoveInDate:  req.MoveInDate,
	}
	if err := dao.DB.Create(&newInterest).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") || strings.Contains(err.Error(), "unique constraint") {
			return errors.New("DUPLICATE_KEY: pet already exists")
		}
		return err
	}
	return nil
}

func (dao *TaskDao) GetApartmentUnitByUnitRentID(unitRentID int) (*model.ApartmentUnit, error) {
	var unit model.ApartmentUnit
	err := dao.DB.Where("unit_rent_id = ?", unitRentID).First(&unit).Error
	if err != nil {
		return nil, err
	}
	return &unit, nil
}

func (dao *TaskDao) GetApartmentBuildingByUnitRentID(unitRentID int) (*model.ApartmentBuilding, error) {
	var unit model.ApartmentUnit
	err := dao.DB.Where("unit_rent_id = ?", unitRentID).First(&unit).Error
	if err != nil {
		return nil, err
	}

	var building model.ApartmentBuilding
	err = dao.DB.Where("company_name = ? AND building_name = ?", unit.CompanyName, unit.BuildingName).First(&building).Error
	if err != nil {
		return nil, err
	}

	return &building, nil
}
func (dao *TaskDao) GetAmenitiesInByUnitRentID(unitRentID int) ([]model.AmenitiesIn, error) {
	var amenities []model.AmenitiesIn
	err := dao.DB.Where("unit_rent_id = ?", unitRentID).Find(&amenities).Error
	if err != nil {
		return nil, err
	}
	return amenities, nil
}
func (dao *TaskDao) GetProvidesByUnitRentID(unitRentID int) ([]model.Provides, error) {
	var unit model.ApartmentUnit
	err := dao.DB.Where("unit_rent_id = ?", unitRentID).First(&unit).Error
	if err != nil {
		return nil, err
	}

	var provides []model.Provides
	err = dao.DB.Where("company_name = ? AND building_name = ?", unit.CompanyName, unit.BuildingName).Find(&provides).Error
	if err != nil {
		return nil, err
	}

	return provides, nil
}
func (dao *TaskDao) CountAvailableUnitsByUnitRentID(unitRentID int) (int, error) {

	var unit model.ApartmentUnit
	err := dao.DB.Where("unit_rent_id = ?", unitRentID).First(&unit).Error
	if err != nil {
		return 0, err
	}
	var count int64
	err = dao.DB.Model(&model.ApartmentUnit{}).
		Where("company_name = ?", unit.CompanyName).
		Where("building_name = ?", unit.BuildingName).
		Where("available_date_for_move_in IS NOT NULL").
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (dao *TaskDao) GetRoomCountsByUnitRentID(unitRentID int) (int, int, int, error) {
	var rooms []model.Rooms
	err := dao.DB.Where("unit_rent_id = ?", unitRentID).Find(&rooms).Error
	if err != nil {
		return 0, 0, 0, err
	}
	bedroomCount := 0
	bathroomCount := 0
	livingRoomCount := 0
	bedroomRegex := regexp.MustCompile(`(?i)bedroom\d*`)
	bathroomRegex := regexp.MustCompile(`(?i)bathroom\d*`)
	livingRoomRegex := regexp.MustCompile(`(?i)livingroom\d*`)

	for _, room := range rooms {
		if bedroomRegex.MatchString(room.Name) {
			bedroomCount++
		} else if bathroomRegex.MatchString(room.Name) {
			bathroomCount++
		} else if livingRoomRegex.MatchString(room.Name) {
			livingRoomCount++
		}
	}

	return bedroomCount, bathroomCount, livingRoomCount, nil
}
