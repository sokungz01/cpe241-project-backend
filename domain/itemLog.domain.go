package domain

import "time"

type ItemLog struct {
	ItemLogID  int       `json:"itemLogID" db:"itemLogID"`
	ItemID     int       `json:"itemID" db:"itemID"`
	ItemQty    int       `json:"qty" db:"qty"`
	StaffID    int       `json:"staffID" db:"staffID"`
	CreateDate time.Time `json:"createDate" db:"createdDate"`
	Staff      User      `json:"staff" db:",prefix=employee."`
	Item       Item      `json:"item" db:",prefix=inventory."`
}

type ItemLogUsecase interface {
	GetAll() (*[]ItemLog, error)
}

type ItemLogRepository interface {
	GetAll() (*[]ItemLog, error)
}
