package model

//"gorm.io/gorm"
//"golang.org/x/crypto/bcrypt"

type User struct {
	Username  string `gorm:"primary_key"`
	FirstName string
	LastName  string
	DOB       string
	Gender    int
	Email     string
	Phone     string
	Passwd    string
}
