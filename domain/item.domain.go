package domain

type Item struct {
	ItemID         int     `json:"itemID" db:"itemID"`
	ItemCategoryID int     `json:"itemCategoryID" db:"itemCategoryID"`
	ItemName       string  `json:"itemName" db:"itemName"`
	ItemCost       float64 `json:"itemCost" db:"itemCost"`
	ItemQty        int     `json:"qty" form:"qty" db:"qty"`
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
