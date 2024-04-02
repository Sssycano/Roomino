package model

//"gorm.io/gorm"
//"golang.org/x/crypto/bcrypt"

type Users struct {
	Username  string `gorm:"primary_key;size:20"`
	FirstName string `gorm:"size:20;not null"`
	LastName  string `gorm:"size:20;not null"`
	DOB       string `gorm:"size:20;not null"`
	Gender    int    `gorm:"type:tinyint;not null"`
	Email     string `gorm:"size:50"`
	Phone     string `gorm:"size:20"`
	Passwd    string `gorm:"size:200"`
	Pets      []Pets `gorm:"foreignKey:Username"`
}

type Pets struct {
	PetName  string `gorm:"primary_key;size:50;not null"`
	PetType  string `gorm:"primary_key;size:50;not null"`
	PetSize  string `gorm:"size:20;not null"`
	Username string `gorm:"primary_key;size:20"`
}
type ApartmentBuilding struct {
	CompanyName   string          `gorm:"size:20;not null;primaryKey"`
	BuildingName  string          `gorm:"size:20;not null;primaryKey"`
	AddrNum       int             `gorm:"not null"`
	AddrStreet    string          `gorm:"size:20;not null"`
	AddrCity      string          `gorm:"size:20;not null"`
	AddrState     string          `gorm:"size:5;not null"`
	AddrZipCode   string          `gorm:"size:5;not null"`
	YearBuilt     int             `gorm:"type:YEAR;not null"`
	ApartmentUnit []ApartmentUnit `gorm:"foreignKey:CompanyName,BuildingName"`
	PetPolicies   []PetPolicy     `gorm:"foreignKey:CompanyName,BuildingName"`
	Provides      []Provides      `gorm:"foreignKey:CompanyName,BuildingName"`
}

type ApartmentUnit struct {
	UnitRentID             int           `gorm:"primaryKey;autoIncrement"`
	CompanyName            string        `gorm:"size:20;not null"`
	BuildingName           string        `gorm:"size:20;not null"`
	UnitNumber             string        `gorm:"size:10;not null"`
	MonthlyRent            int           `gorm:"not null"`
	SquareFootage          int           `gorm:"not null"`
	AvailableDateForMoveIn string        `gorm:"type:date;not null"`
	Rooms                  []Rooms       `gorm:"foreignKey:UnitRentID"`
	AmenitiesIn            []AmenitiesIn `gorm:"foreignKey:UnitRentID"`
}
type Rooms struct {
	Name          string `gorm:"size:20;not null;primaryKey"`
	SquareFootage int    `gorm:"not null"`
	Description   string `gorm:"size:50;not null"`
	UnitRentID    uint   `gorm:"not null;primaryKey"`
}
type PetPolicy struct {
	CompanyName     string `gorm:"primaryKey;size:20"`
	BuildingName    string `gorm:"primaryKey;size:20"`
	PetType         string `gorm:"primaryKey;size:50"`
	PetSize         string `gorm:"primaryKey;size:20"`
	IsAllowed       bool   `gorm:"not null"`
	RegistrationFee int
	MonthlyFee      int
}

type Amenities struct {
	AType       string        `gorm:"primaryKey;size:20"`
	Description string        `gorm:"size:100;not null"`
	AmenitiesIn []AmenitiesIn `gorm:"foreignKey:AType"`
	Provides    []Provides    `gorm:"foreignKey:AType"`
}

type AmenitiesIn struct {
	AType      string `gorm:"primaryKey;size:20;not null"`
	UnitRentID uint   `gorm:"primaryKey;autoIncrement;not null"`
}

type Provides struct {
	AType        string `gorm:"primaryKey;size:20"`
	CompanyName  string `gorm:"primaryKey;size:20;not null"`
	BuildingName string `gorm:"primaryKey;size:20;not null"`
	Fee          int    `gorm:"not null"`
	WaitingList  int    `gorm:"not null"`
}

func (Users) TableName() string {
	return "Users"
}
func (Pets) TableName() string {
	return "Pets"
}
func (Rooms) TableName() string {
	return "Rooms"
}
func (Amenities) TableName() string {
	return "Amenities"
}
func (Provides) TableName() string {
	return "Provides"
}
func (ApartmentBuilding) TableName() string {
	return "ApartmentBuilding"
}
func (ApartmentUnit) TableName() string {
	return "ApartmentUnit"
}
func (PetPolicy) TableName() string {
	return "PetPolicy"
}
func (AmenitiesIn) TableName() string {
	return "AmenitiesIn"
}
