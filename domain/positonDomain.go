package domain

type Position struct {
	PositionID     int     `json:"positionID" db:"positionID"`
	PositionName   string  `json:"positionName" db:"positionName"`
	PositionSalary float64 `json:"positionSalary" db:"positionSalary"`
}

type PositionRepository interface {
	Create(position *Position) error
	FindByPositionName(positionName string) (*Position, error)
	GetPositionByID(id int) (*Position, error)
	UpdatePosition(id int, newPositionData *Position) error
	DeletePosition(id int) error
	GetAll() (*[]Position, error)
}

type PositionUsecase interface {
	Create(position *Position) error
	FindByPositionName(positionName string) (*Position, error)
	GetPositionByID(id int) (*Position, error)
	UpdatePosition(id int, newPositionData *Position) error
	DeletePosition(id int) error
	GetAll() (*[]Position, error)
}
