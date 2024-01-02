package models

import (
	"strings"

	"github.com/rralcala/martyn/lib/log"
	"gorm.io/gorm/clause"
)

type AccountModel struct{}

type Account struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
}

func (*AccountModel) GetSingleItem(key int64) (*Account, error) {

	var item Account
	if err := db.Model(item).Where("id = ?", key).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (*AccountModel) TotalCount() int64 {
	var count int64
	db.Model(&Account{}).Count(&count)
	return count
}

func (*AccountModel) Create(newItem *Account) {
	db.Create(newItem)
}

func (*AccountModel) Delete(id []*Account) {
	for _, i := range id {
		db.Delete(i)
	}
}

func (*AccountModel) GetList(sort []string, itemRange []int, filters map[string]interface{}) []Account {
	items := []Account{}
	db := db.Preload(clause.Associations)
	if len(sort) == 2 {
		log.Info("Will sort")
		db = db.Order(strings.ToLower(strings.Join(sort[:], " ")))
	} else {
		db = db.Order("description")
	}

	if len(itemRange) == 2 {
		log.Info("Limited range")
		db = db.Offset(itemRange[0]).Limit(itemRange[1])
	}
	if len(filters) > 0 {
		db = db.Where(filters)
	}
	db.Find(&items)
	return items
}
