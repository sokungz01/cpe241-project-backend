package domain

type UserPosition struct{
	PositionID		int		`json:"positionid" db:"positionid"`
	PositionName	string	`json:"positionname" db:"positionname"`
	PositionSalary	float64  `json:"positionsalary" db:"positionsalary"`
}

type PositionRepository interface{
	Create (position *UserPosition) error
}

type PositionUsecase interface{
	Create (position *UserPosition) error
}