package domain

type MachineType struct {
	MachineTypeID   int    `json:"machinetypeID" db:"machineTypeID"`
	MachineTypeName string `json:"machinetypeName" db:"machineTypeName"`
}

type MachineTypeRepository interface {
	CreateMachineType(mtype MachineType) error
	GetAllMachineType() (*[]MachineType, error)
	GetOneMachineTypeByName(typeName string) (*MachineType, error)
	GetOneMachineTypeByID(id int) (*MachineType, error)
	UpDateMachineType(id int, newData MachineType) (*MachineType, error)
	DeleteMachineType(id int) error
}

type MachineTypeUsecase interface {
	CreateMachineType(mtype MachineType) error
	GetAllMachineType() (*[]MachineType, error)
	GetOneMachineTypeByName(typeName string) (*MachineType, error)
	GetOneMachineTypeByID(id int) (*MachineType, error)
	UpDateMachineType(id int, newData MachineType) (*MachineType, error)
	DeleteMachineType(id int) error
}
