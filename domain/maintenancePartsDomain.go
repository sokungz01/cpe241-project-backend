package domain

import (
	"time"
)

type MaintenanceParts struct {
	MaintenancePartID int       `json:"maintenancePart" db:"maintenancePart"`
	ServiceID         int       `json:"serviceID" db:"serviceID"`
	ItemID            int       `json:"itemID" db:"itemID"`
	Qty               int       `json:"qty" db:"qty"`
	CreatedDate       time.Time `json:"createdDate" db:"createdDate"`
}

type MaintenancePartsRepository interface {
	CreateMaintenanceParts(serviceID int, itemID int, qty int, createdDate time.Time) error
}
