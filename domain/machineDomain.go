package domain

import (
	"time"
)

type Machine struct {
	MachineID     int       `json:"machineid" db:"machineid"`
	MachineName   string    `json:"machinename" db:"machinename"`
	MachineBrand  string    `json:"machinebrand" db:"machinebrand"`
	MachineTypeID int       `json:"machinetypeid" db:"machinetypeid"`
	StartDate     time.Time `json:"startdate" db:"startdate"`
	EndDate       time.Time `json:"enddate" db:"enddate"`
	Description   string    `json:"desc" db:"desciption"`
	Status        bool      `json:"status" db:"status"`
}