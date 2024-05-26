package domain

type Item struct {
	ItemID         int     `json:"itemID,omitempty" db:"itemID"`
	ItemCategoryID int     `json:"itemCategoryID,omitempty" db:"itemCategoryID"`
	ItemName       string  `json:"itemName,omitempty" db:"itemName"`
	ItemCost       float64 `json:"itemCost,omitempty" db:"itemCost"`
	ItemQty        int     `json:"qty,omitempty" form:"qty" db:"qty"`
	StaffID        int     `json:"staffID,omitempty" db:"staffID"`
	IsDelete       int     `json:"isDelete,omitempty" db:"isDelete"`
}

type ItemUseCase interface {
	CreateItem(item *Item) (*Item, error)
	GetAllItem() (*[]Item, error)
	FindByID(id int) (*Item, error)
	UpdateItem(id int, item *Item) (*Item, error)
	DeleteItem(id int) error
}

type ItemRepository interface {
	CreateItem(item *Item) (*Item, error)
	GetAllItem() (*[]Item, error)
	FindByID(id int) (*Item, error)
	UpdateItem(id int, item *Item) (*Item, error)
	DeleteItem(id int) error
}
