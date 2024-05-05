package domain

type Position struct{
	PositionID		int		`json:"positionid" db:"positionID"`
	PositionName	string	`json:"positionname" db:"positionName"`
	PositionSalary	float64  `json:"positionsalary" db:"positionSalary"`
}

type PositionRepository interface{
	Create(position *Position) error
	FindByPositionName(positionName string) (*Position,error)
	GetAll()(*[]Position,error)
}

type PositionUsecase interface{
	Create(position *Position) error
	FindByPositionName(positionName string)(*Position,error)
	GetAll() (*[]Position,error)
}