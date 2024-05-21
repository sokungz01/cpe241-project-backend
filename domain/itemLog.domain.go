package domain

type ItemLog struct {
	ItemLogID  int `json:"itemLogID" db:"itemLogID"`
	ItemID     int `json:"itemID" db:"itemID"`
	ItemQty    int `json:"qty" db:"qty"`
	StaffID    int `json:"staffID" db:"staffID"`
	CreateDate int `json:"createDate" db:"createDate"`
}

type ItemLogUsecase interface {
	GetAll() (*[]ItemLog, error)
}

type ItemLogRepository interface {
	GetAll() (*[]ItemLog, error)
}
