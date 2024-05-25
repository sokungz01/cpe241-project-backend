package usecase

import (
	"errors"

	"github.com/sokungz01/cpe241-project-backend/domain"
)

type itemLogUsecase struct {
	repository     domain.ItemLogRepository
	itemRepository domain.ItemRepository
	userRepository domain.UserRepository
}

func NewItemLogUsecase(repository domain.ItemLogRepository, itemRepository domain.ItemRepository, userRepository domain.UserRepository) domain.ItemLogUsecase {
	return &itemLogUsecase{
		repository:     repository,
		itemRepository: itemRepository,
		userRepository: userRepository,
	}
}
func (u *itemLogUsecase) GetAll() (*[]domain.ItemLog, error) {
	response, err := u.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *itemLogUsecase) CreateItemLog(itemLog *domain.ItemLog) (*domain.ItemLog, error) {
	query, err := u.itemRepository.FindByID(itemLog.ItemID)
	if err != nil {
		return nil, err
	}
	_, err = u.userRepository.GetById(itemLog.StaffID)
	if err != nil {
		return nil, err
	}
	item := new(domain.Item)
	item = query
	if item.ItemQty+itemLog.ItemQty < 0 {
		return nil, errors.New("resulting item quantity cannot be negative")
	}
	if itemLog.IsAdd {
		item.ItemQty += itemLog.ItemQty
	} else {
		item.ItemQty -= itemLog.ItemQty
	}
	_, err = u.itemRepository.UpdateItem(itemLog.ItemID, item)
	if err != nil {
		return nil, err
	}

	response, err := u.repository.CreateItemLog(itemLog)
	if err != nil {
		return nil, err
	}

	return response, nil
}
