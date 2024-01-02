package models

import (
	"strings"

	"github.com/rralcala/martyn/lib/log"
	"gorm.io/gorm/clause"
)

type ProviderModel struct{}

type Provider struct {
	ID   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (*ProviderModel) GetSingleItem(key int64) (*Provider, error) {

	var item Provider
	if err := db.Model(item).Where("id = ?", key).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (*ProviderModel) TotalCount() int64 {
	var count int64
	db.Model(&Provider{}).Count(&count)
	return count
}

func (*ProviderModel) Create(newItem *Provider) {
	db.Create(newItem)
}

func (*ProviderModel) Delete(id []*Provider) {
	for _, i := range id {
		db.Delete(i)
	}
}

func (*ProviderModel) GetList(sort []string, itemRange []int, filters map[string]interface{}) []Provider {
	items := []Provider{}
	db := db.Preload(clause.Associations)
	if len(sort) == 2 {
		log.Info("Will sort")
		db = db.Order(strings.ToLower(strings.Join(sort[:], " ")))
	} else {
		db = db.Order("name")
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
