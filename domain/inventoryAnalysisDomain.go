package domain

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type InventoryAnalysis struct {
	Date         time.Time `json:"Date" db:"Date"`
	CategoryName string    `json:"categoryName" db:"categoryName"`
	TotalAdded   int       `json:"totalAdded" db:"totalAdded"`
	TotalRemoved int       `json:"totalRemoved" db:"totalRemoved"`
}

type MachineTypeErrorAnalysis struct {
	MachineTypeName      string    `json:"machineTypeName" db:"machineTypeName"`
	TotalServiceRequests int       `json:"totalServiceRequests" db:"totalServiceRequests"`
	RequestDate          time.Time `json:"requestDate" db:"requestDate"`
}

type EmployeeEngagementAnalysis struct {
	EmployeeID          int    `json:"employeeID" db:"employeeID"`
	Name                string `json:"name" db:"name"`
	Surname             string `json:"surname" db:"surname"`
	MaintenanceCount    int    `json:"maintenanceCount" db:"maintenanceCount"`
	ToltalInventoryUsed int    `json:"totalInventoryUsed" db:"totalInventoryUsed"`
	InventoryItemsUsed  string `json:"inventoryItemsUsed" db:"inventoryItemsUsed"`
}

type MaintenanceCostAnalysis struct {
	ErrorName            string  `json:"errorName" db:"errorName"`
	ErrorCount           int     `json:"errorCount" db:"errorCount"`
	TotalmaintenanceCost float64 `json:"totalMaintenanceCost" db:"totalMaintenanceCost"`
}

type AnalysisRepository interface {
	GetInventoryAnanlysis() (*[]InventoryAnalysis, error)
	GetMachineTypeErrorAnalysis() (*[]MachineTypeErrorAnalysis, error)
	GetEmployeeEngagementAnalysis() (*[]EmployeeEngagementAnalysis, error)
	GetMaintenanceCostAnalysis() (*[]MaintenanceCostAnalysis, error)
}

type InventoryAnalysisController interface {
	GetInventoryAnanlysis(c *fiber.Ctx) error
	GetMachineTypeErrorAnalysis(c *fiber.Ctx) error
	GetEmployeeEngagementAnalysis(c *fiber.Ctx) error
	GetMaintenanceCostAnalysis(c *fiber.Ctx) error
}
