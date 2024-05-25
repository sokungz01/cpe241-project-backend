package repository

import (
	"time"

	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type maintenancePartRepository struct {
	db *platform.Mysql
}

func NewMaintenancePartsRepository(db *platform.Mysql) domain.MaintenancePartsRepository {
	return &maintenancePartRepository{db: db}
}

func (r *maintenancePartRepository) CreateMaintenanceParts(serviceID int, itemID int, qty int, createdDate time.Time) error {
	_, err := r.db.Exec("INSERT INTO `maintenanceParts` (`serviceID`,`itemID`,`qty`,`createdDate`)"+
		" VALUE (?,?,?,?)", serviceID, itemID, qty, createdDate)
	return err
}
