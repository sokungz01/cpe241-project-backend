package domain

type ErrorType struct {
	ErrorTypeID int    `json:"errorTypeID" db:"errorTypeID"`
	ErrorName   string `json:"errorName" db:"errorName"`
}

type ErrorTypeUseCase interface {
	CreateErrorType(elem *ErrorType) (*ErrorType, error)
	GetAllErrorType() (*[]ErrorType, error)
	FindByID(id int) (*ErrorType, error)
	UpdateErrorType(id int, elem *ErrorType) (*ErrorType, error)
	DeleteErrorType(id int) error
}

type ErrorTypeRepository interface {
	CreateErrorType(elem *ErrorType) (*ErrorType, error)
	GetAllErrorType() (*[]ErrorType, error)
	FindByID(id int) (*ErrorType, error)
	UpdateErrorType(id int, elem *ErrorType) (*ErrorType, error)
	DeleteErrorType(id int) error
}
