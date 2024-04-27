package types

type UserServiceReq struct {
	Username  string `json:"username" binding:"required,max=20"`
	FirstName string `json:"first_name" binding:"required,max=20"`
	LastName  string `json:"last_name" binding:"required,max=20"`
	DOB       string `json:"dob" binding:"required,max=20"`
	Gender    int    `json:"gender" binding:"required"`
	Email     string `json:"email" binding:"omitempty,email,max=50"`
	Phone     string `json:"phone" binding:"omitempty,max=20"`
	Passwd    string `json:"passwd" binding:"omitempty,max=200"`
}
