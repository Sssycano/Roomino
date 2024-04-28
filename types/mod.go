package types

type UserServiceReq struct {
	Username  string `json:"username" binding:"required,max=20"`
	FirstName string `json:"first_name" binding:"omitempty,max=20"`
	LastName  string `json:"last_name" binding:"omitempty,max=20"`
	DOB       string `json:"dob" binding:"omitempty,max=20"`
	Gender    int    `json:"gender" binding:"omitempty"`
	Email     string `json:"email" binding:"omitempty,email,max=50"`
	Phone     string `json:"phone" binding:"omitempty,max=20"`
	Passwd    string `json:"passwd" binding:"omitempty,max=200"`
}
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type UserResp struct {
	UserName string `json:"user_name" form:"user_name" example:"FanOne"`
	//CreateAt int64  `json:"create_at" form:"create_at"`
}

type UnitInforeq struct {
	CompanyName  string `json:"company_name" binding:"required,max=20"`
	BuildingName string `json:"building_name" binding:"required,max=20"`
	//Username     string `json:"username" binding:"required,max=20"`
}
type UnitInfoResp struct {
	UnitRentID             int    `json:"unit_rent_id" form:"unit_rent_id"`
	MonthlyRent            int    `json:"monthly_rent" form:"monthly_rent"`
	SquareFootage          int    `json:"square_footage" form:"square_footage"`
	AvailableDateForMoveIn string `json:"available_date_for_move_in" form:"available_date_for_move_in" binding:"required"`
	IsPetAllowed           bool   `json:"is_pet_allowed" form:"is_pet_allowed"`
}
