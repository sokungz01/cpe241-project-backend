package domain

import (
	"database/sql"
	"time"
)

type Machine struct {
	MachineID     int           `json:"machineID" db:"machineID"`
	MachineName   string        `json:"machineName" db:"machineName"`
	MachineBrand  string        `json:"machineBrand" db:"machineBrand"`
	MachineTypeID int           `json:"machineTypeID" db:"machineTypeID"`
	StartDate     time.Time     `json:"startDate" db:"startDate"`
	EndDate       *sql.NullTime `json:"endDate" db:"endDate"`
	Description   string        `json:"desc" db:"description"`
	Status        int           `json:"status" db:"status"`
}

type MachineRepository interface {
	CreateMachine(newMachine *Machine) (*Machine, error)
	GetAllMachine() (*[]Machine, error)
	GetMachineByID(id int) (*Machine, error)
	GetMachineByName(machineName string) (*[]Machine, error)
	UpdateMachineData(id int, newMachineData *Machine) error
	DeleteMachine(id int) error
	UpdateMachineStatus(id int, status int) error
}

type MachineUsecase interface {
	CreateMachine(newMachine *Machine) (*Machine, error)
	GetAllMachine() (*[]Machine, error)
	GetMachineByID(id int) (*Machine, error)
	GetMachineByName(machineName string) (*[]Machine, error)
	UpdateMachineData(id int, newMachineData *Machine) (*Machine, error)
	DeleteMachine(id int) error
	UpdateMachineStatus(id int, status int) error
}
