package db

import (
	
	"github.com/subhashsras/backend_service/models"
	"gorm.io/gorm"
)

type ItemRepo interface{
	GetItems() ([]models.Item, error)
	GetItem(itemId uint64) (models.Item, error)
	CreateItem(item models.Item) (error)
}

type itemRepo struct{
	db *gorm.DB
}

func NewItemRepo() *itemRepo {
	return &itemRepo{
		db: DBClient,
	}
}


func (i *itemRepo) GetItems() ([]models.Item, error) {
	var items []models.Item

	result := i.db.Find(&items)
	if result.Error != nil {
		return items, result.Error
	}

	return items, nil
}

func (i *itemRepo) GetItem(itemId uint64) (models.Item, error) {
	var item models.Item

	result := i.db.Where("id = ?", itemId).Find(&item)
	if result.Error != nil {
		return item, result.Error
	}

	return item, nil
}

func (i *itemRepo) CreateItem(item models.Item)  error {
	
	result := i.db.Create(&item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}


