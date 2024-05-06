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

type MachineType struct {
	MachineTypeID   int    `json:"machinetypeid" db:"machineTypeID"`
	MachineTypeName string `json:"machinetypename" db:"machineTypeName"`
}

type MachineRepository interface {
	GetAll() (*[]Machine, error)

	CreateMachineType(mtype MachineType) error
	GetOneMachineTypeByName(typeName string) (*MachineType, error)
	GetOneMachineTypeByID(id int) (*MachineType, error)
	UpDateMachineType(id int, newData MachineType) (*MachineType, error)
	DeleteMachineType(id int) error
}

type MachineUseCase interface {
	GetAll() (*[]Machine, error)
	CreateMachineType(mtype MachineType) error
	GetOneMachineTypeByName(typeName string) (*MachineType, error)
	UpDateMachineType(id int, newData MachineType) (*MachineType, error)
	DeleteMachineType(id int) error
}
