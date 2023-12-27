package models

type Account struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
}

type Provider struct {
	ID   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func FindProvider(key int64) *Provider {

	var item Provider
	if err := DB.Model(item).Where("id = ?", key).First(&item).Error; err != nil {
		return nil
	}
	return &item
}

func FindAccount(key int64) *Account {

	var item Account
	if err := DB.Model(item).Where("id = ?", key).First(&item).Error; err != nil {
		return nil
	}
	return &item
}

type CostCenter struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
}

func FindCostCenter(key int64) *CostCenter {

	var item CostCenter
	if err := DB.Model(item).Where("id = ?", key).First(&item).Error; err != nil {
		return nil
	}
	return &item
}
