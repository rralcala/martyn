package models

import (
	"strings"

	"github.com/rralcala/martyn/lib/log"
	"gorm.io/gorm/clause"
)

type CostCenterModel struct{}

type CostCenter struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
}

func (*CostCenterModel) GetSingleItem(key int64) (*CostCenter, error) {

	var item CostCenter
	if err := db.Model(item).Where("id = ?", key).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (*CostCenterModel) TotalCount() int64 {
	var count int64
	db.Model(&CostCenter{}).Count(&count)
	return count
}

func (*CostCenterModel) Create(newItem *CostCenter) {
	db.Create(newItem)
}

func (*CostCenterModel) Delete(id []*CostCenter) {
	for _, i := range id {
		db.Delete(i)
	}
}

func (*CostCenterModel) GetList(sort []string, itemRange []int, filters map[string]interface{}) []CostCenter {
	var items []CostCenter
	db := db.Preload(clause.Associations)
	if len(sort) == 2 {
		log.Info("Will sort")
		db = db.Order(strings.ToLower(strings.Join(sort[:], " ")))
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
