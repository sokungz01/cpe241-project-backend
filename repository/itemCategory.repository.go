package repository

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type ItemCaterepository struct {
	db *platform.Mysql
}

func NewItemCategoryRepository(db *platform.Mysql) domain.ItemCategoryRepository {
	return &ItemCaterepository{db: db}
}

func (r *ItemCaterepository) CreateItemCategory(category *domain.ItemCategory) (*domain.ItemCategory, error) {
	_, err := r.db.NamedExec("INSERT INTO `itemCategory` (`categoryName`)"+
		"VALUE (:categoryName)", category)

	if err != nil {
		return nil, err
	}

	response := new(domain.ItemCategory)
	_ = r.db.Get(response, "SELECT * FROM `itemCategory` WHERE categoryID IN (SELECT LAST_INSERT_ID() as id)")

	return response, nil
}

func (r *ItemCaterepository) GetAllItemCategory() (*[]domain.ItemCategory, error) {
	response := make([]domain.ItemCategory, 0)
	if err := r.db.Select(&response, "SELECT * FROM `itemCategory`"); err != nil {
		return nil, err
	}
	return &response, nil
}

func (r *ItemCaterepository) FindByID(id int) (*domain.ItemCategory, error) {
	response := new(domain.ItemCategory)
	err := r.db.Get(response, "SELECT *"+
		"FROM `itemCategory`"+
		"WHERE `categoryID` = ?", id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *ItemCaterepository) UpdateItemCategory(id int, category *domain.ItemCategory) (*domain.ItemCategory, error) {
	_, err := r.db.Exec("UPDATE `itemCategory`"+
		"SET `categoryName`= ? WHERE `categoryID`= ?", category.CategoryName, id)
	if err != nil {
		return nil, err
	}
	response, _ := r.FindByID(id)
	return response, nil
}

func (r *ItemCaterepository) DeleteItemCategory(id int) error {
	_, err := r.db.Exec("DELETE FROM `itemCategory` WHERE `categoryID` = ?", id)
	if err != nil {
		return err
	}
	return nil
}
