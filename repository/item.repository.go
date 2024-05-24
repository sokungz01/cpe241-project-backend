package repository

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type itemRepository struct {
	db *platform.Mysql
}

func NewItemRepository(db *platform.Mysql) domain.ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) CreateItem(item *domain.Item) (*domain.Item, error) {
	_, err := r.db.NamedExec("INSERT INTO `inventory` (`itemCategoryID`, `itemName`, `itemCost`, `qty`)"+
		"VALUE (:itemCategoryID, :itemName, :itemCost, :qty)", item)
	if err != nil {
		return nil, err
	}
	response := new(domain.Item)
	_ = r.db.Get(response, "SELECT * FROM `inventory` WHERE itemID IN (SELECT LAST_INSERT_ID() as id)")
	return response, nil
}

func (r *itemRepository) GetAllItem() (*[]domain.Item, error) {
	response := make([]domain.Item, 0)
	if err := r.db.Select(&response, "SELECT * FROM `inventory`"); err != nil {
		return nil, err
	}
	return &response, nil
}

func (r *itemRepository) FindByID(id int) (*domain.Item, error) {
	response := new(domain.Item)
	err := r.db.Get(response, "SELECT *"+
		"FROM `inventory`"+
		"WHERE `itemID` = ?", id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *itemRepository) UpdateItem(id int, item *domain.Item) (*domain.Item, error) {
	_, err := r.db.Exec("UPDATE `inventory`"+
		"SET `itemName`= ? , `itemCategoryID` = ?, `itemCost` = ? , `qty` = ? WHERE `itemID`= ?", item.ItemName, item.ItemCategoryID, item.ItemCost, item.ItemQty, id)
	if err != nil {
		return nil, err
	}
	response, _ := r.FindByID(id)
	return response, nil
}

func (r *itemRepository) DeleteItem(id int) error {
	_, err := r.db.Exec("DELETE FROM `inventory` WHERE `itemID` = ?", id)

	if err != nil {
		return err
	}

	return nil
}
