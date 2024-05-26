package repository

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type analysisRepo struct {
	db *platform.Mysql
}

func NewAnalysisRepo(db *platform.Mysql) domain.AnalysisRepository {
	return &analysisRepo{db: db}
}

func (r *analysisRepo) GetInventoryAnanlysis() (*[]domain.InventoryAnalysis, error) {
	response := new([]domain.InventoryAnalysis)
	err := r.db.Select(response, "SELECT DATE(il.createdDate) AS `Date`,`ic`.`categoryName` AS `categoryName`,"+
		"SUM(CASE WHEN il.isAdd = 1 THEN il.qty ELSE 0 END) AS `totalAdded`,SUM(CASE WHEN il.isAdd = 0 THEN il.qty ELSE 0 END) AS `totalRemoved` "+
		"FROM `inventory` AS `i` "+
		"INNER JOIN `itemCategory` AS `ic` ON `i`.`itemCategoryID` = `ic`.`categoryID` "+
		"INNER JOIN `itemLog` AS `il` ON `i`.`itemID` = `il`.`itemID` "+
		"GROUP BY `Date`,`ic`.`categoryName`")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *analysisRepo) GetMachineTypeErrorAnalysis() (*[]domain.MachineTypeErrorAnalysis, error) {
	response := new([]domain.MachineTypeErrorAnalysis)
	err := r.db.Select(response,
		"SELECT `mt`.`machineTypeName` AS `machineTypeName`,"+
			"COUNT(`sr`.`serviceID`) AS `totalServiceRequests`,"+
			"DATE(`sr`.`CreatedDate`) AS `requestDate`"+
			"FROM `serviceRequest` AS `sr`"+
			"INNER JOIN `machine` AS `m` ON `sr`.`machineID` = `m`.`machineID` "+
			"INNER JOIN `machineType` AS `mt` ON `m`.`machineTypeID` = `mt`.`machineTypeID` "+
			"GROUP BY DATE(`sr`.`createdDate`), `mt`.`machineTypeName` "+
			"ORDER BY DATE(`sr`.`createdDate`)")

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *analysisRepo) GetEmployeeEngagementAnalysis() (*[]domain.EmployeeEngagementAnalysis, error) {
	response := new([]domain.EmployeeEngagementAnalysis)
	err := r.db.Select(response,
		"SELECT "+
			"`e`.`employeeID` AS `employeeID`,"+
			"`e`.`name` AS `name`,"+
			"`e`.`surname` AS `surname`,"+
			"COUNT(`sres`.`staffServiceID`) AS `maintenanceCount`,"+
			"SUM(`mp`.`qty`) AS `totalInventoryUsed`,"+
			"GROUP_CONCAT(i.itemName SEPARATOR ', ') AS `inventoryItemsUsed` "+
			"FROM `serviceResponse` AS `sres`"+
			"INNER JOIN `employee` AS `e` ON `sres`.`staffID` = `e`.`employeeID` "+
			"INNER JOIN `maintenanceParts` AS `mp` ON `sres`.`staffServiceID` = `mp`.`serviceID` "+
			"INNER JOIN `inventory` AS `i` ON `mp`.`itemID` = `i`.`itemID` "+
			"GROUP BY `e`.`employeeID`, `e`.`name`, `e`.`surname`")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *analysisRepo) GetMaintenanceCostAnalysis() (*[]domain.MaintenanceCostAnalysis, error) {
	response := new([]domain.MaintenanceCostAnalysis)
	err := r.db.Select(response,
		"SELECT "+
			"`et`.`errorName` AS `errorName`,"+
			"COUNT(el.errorTypeID) AS `errorCount`,"+
			"SUM(mp.qty * i.itemCost) AS `totalMaintenanceCost` "+
			"FROM `errorLog` AS `el` "+
			"INNER JOIN `errorType` AS `et` ON `et`.`errorTypeID` = `el`.`errorTypeID` "+
			"INNER JOIN `serviceResponse` AS `sres` ON `sres`.`requestedServiceID` = `el`.`serviceID` "+
			"INNER JOIN `maintenanceParts` AS `mp` ON `sres`.`staffServiceID` = `mp`.`serviceID` "+
			"INNER JOIN `inventory` AS `i` ON `mp`.`itemID` = `i`.`itemID` "+
			"GROUP BY `el`.`errorTypeID`")
	if err != nil {
		return nil, err
	}
	return response, nil
}
