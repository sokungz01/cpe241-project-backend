package domain

type ItemCategory struct {
	CategoryID   int    `json:"categoryID" db:"categoryID"`
	CategoryName string `json:"categoryName" db:"categoryName"`
}

type ItemCategoryUseCase interface {
	CreateItemCategory(category *ItemCategory) (*ItemCategory, error)
	GetAllItemCategory() (*[]ItemCategory, error)
	FindByID(id int) (*ItemCategory, error)
	UpdateItemCategory(id int, category *ItemCategory) (*ItemCategory, error)
	DeleteItemCategory(id int) error
}

type ItemCategoryRepository interface {
	CreateItemCategory(category *ItemCategory) (*ItemCategory, error)
	GetAllItemCategory() (*[]ItemCategory, error)
	FindByID(id int) (*ItemCategory, error)
	UpdateItemCategory(id int, category *ItemCategory) (*ItemCategory, error)
	DeleteItemCategory(id int) error
}
