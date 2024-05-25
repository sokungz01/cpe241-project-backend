package domain

import (
	"database/sql"
	"time"
)

type ErrorLog struct {
	ErrorID          int          `json:"errorID" db:"errorID"`
	ErrorTypeID      int          `json:"errorTypeID" db:"errorTypeID"`
	ServiceID        int          `json:"serviceID" db:"serviceID"`
	ErrorDescription string       `json:"errorDescription" db:"errorDescription"`
	CreatedDate      time.Time    `json:"createdDate" db:"createdDate"`
	UpdateDate       sql.NullTime `json:"updateDate" db:"updateDate"`
}

type ErrorlogRepository interface {
	Create(newError *ErrorLog) (*ErrorLog, error)
	FindByServiceID(id int) (*[]ErrorLog, error)
}
