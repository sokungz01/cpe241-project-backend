package usecase

import (
	"errors"

	"github.com/sokungz01/cpe241-project-backend/domain"
)

type itemUsecase struct {
	itemRepository domain.ItemRepository
	log            domain.ItemLogUsecase
}

func NewItemUsecase(itemRepository domain.ItemRepository, log domain.ItemLogUsecase) domain.ItemUseCase {
	return &itemUsecase{itemRepository: itemRepository,
		log: log}
}

func (u *itemUsecase) CreateItem(item *domain.Item) (*domain.Item, error) {
	if item.ItemName == "" || item.ItemCost == 0 || item.ItemQty == 0 || item.ItemCategoryID == 0 {
		return nil, errors.New("erorr! body empty")
	}
	dummy := new(domain.Item)
	adding := new(domain.ItemLog)

	dummy.ItemCategoryID = item.ItemCategoryID
	dummy.ItemName = item.ItemName
	dummy.ItemCost = item.ItemCost
	dummy.ItemQty = 0
	dummy.StaffID = item.StaffID

	response, err := u.itemRepository.CreateItem(dummy)
	if err != nil {
		return nil, errors.New("erorr! cannot create new item ")
	}
	adding.ItemID = response.ItemID
	adding.ItemQty = item.ItemQty
	adding.StaffID = item.StaffID
	adding.IsAdd = true
	_, logerr := u.log.CreateItemLog(adding)
	if logerr != nil {
		return nil, errors.New("erorr! cannot create new item due to logging error ")
	}
	response.ItemQty = item.ItemQty
	return response, nil
}

func (u *itemUsecase) GetAllItem() (*[]domain.Item, error) {
	response, err := u.itemRepository.GetAllItem()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *itemUsecase) FindByID(id int) (*domain.Item, error) {
	response, err := u.itemRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *itemUsecase) UpdateItem(id int, item *domain.Item) (*domain.Item, error) {
	if item.ItemName == "" || id == 0 {
		return nil, errors.New("erorr! body empty")
	}
	response, err := u.itemRepository.UpdateItem(id, item)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *itemUsecase) DeleteItem(id int) error {
	err := u.itemRepository.DeleteItem(id)
	if err != nil {
		return err
	}
	return nil
}
