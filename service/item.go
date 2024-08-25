package service

import (
	"log"

	"github.com/subhashsras/backend_service/db"
	"github.com/subhashsras/backend_service/models"
)

type ItemService interface {
	GetItems() ([]models.Item, error)
	GetItem(itemId uint64) (models.Item, error)
	CreateItem(item models.Item) (error)
}

type itemService struct {
	itemRepo db.ItemRepo
}

func NewItemService(itemRepo db.ItemRepo) *itemService {
	return &itemService{
		itemRepo: itemRepo,
	}
}

type CustomValidationError struct {
	msg string
}

func (c *CustomValidationError) Error() string {
	return c.msg
}

func (i *itemService) GetItems() ([]models.Item, error) {
	items, err := i.itemRepo.GetItems()
	if err != nil {
		log.Println("Error while fetching items:", err.Error())
		return items, err
	}

	return items, nil

}

func (i *itemService) GetItem(itemId uint64) (models.Item, error) {
	var item models.Item
	if itemId <= 0 {
		err := &CustomValidationError{msg: "Item Id should be greater than 0"}
		log.Println(err.Error())
		return item, err
	}
	item, err := i.itemRepo.GetItem(itemId)
	if err != nil {
		log.Println("Error while fetching items by item id:", err.Error())
		return item, err
	}

	return item, nil

}

func (i *itemService) CreateItem(item models.Item) error {
	err := i.itemRepo.CreateItem(item)
	if err != nil {
		log.Println("Error while creating item:", err.Error())
		return err
	}

	return nil

}
