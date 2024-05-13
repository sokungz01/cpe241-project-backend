package domain

import "time"

type Machine struct {
	MachineID     int       `json:"machineid" db:"machineID"`
	MachineName   string    `json:"machinename" db:"machineName"`
	MachineBrand  string    `json:"machinebrand" db:"machineBrand"`
	MachineTypeID int       `json:"machinetypeid" db:"machineTypeID"`
	StartDate     time.Time `json:"startdate" db:"startDate"`
	EndDate       time.Time `json:"enddate" db:"endDate"`
	Description   string    `json:"desc" db:"description"`
	Status        int       `json:"status" db:"status"`
}

type MachineRepository interface {
	CreateMachine(newMachine *Machine) (*Machine, error)
	GetAllMachine()(*[]Machine,error)
	GetMachineByID(id int) (*Machine,error)
}

type MachineUsecase interface {
	CreateMachine(newMachine *Machine) (*Machine, error)
	GetAllMachine()(*[]Machine,error)
	GetMachineByID(id int) (*Machine,error)
}
