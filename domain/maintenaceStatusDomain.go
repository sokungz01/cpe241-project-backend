package domain

type MaintenanceStatus struct {
	StatusID   int    `json:"statusID" db:"statusID"`
	StatusName string `json:"statusName" db:"statusName"`
}

type MaintenanceStatusRepo interface {
	GetAll() (*[]MaintenanceStatus, error)
}

type MaintenanceStatusUsecase interface {
	GetAll() (*[]MaintenanceStatus, error)
}
