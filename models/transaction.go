package models

import (
	"strings"

	"github.com/rralcala/martyn/lib/log"
	"gorm.io/gorm/clause"
)

type TransactionModel struct{}

// album represents data about a record album.
type Transaction struct {
	ID           int64  `json:"id"`
	Date         string `json:"date"`
	ProviderID   int64
	Provider     *Provider `json:"provider"`
	Description  string    `json:"description"`
	Amount       int64     `json:"amount"`
	CostCenterID int64
	CostCenter   *CostCenter `json:"cost_center"`
	AccountID    int64
	Account      *Account `json:"account"`
}

// album represents data about a record album.
type TransactionOutput struct {
	ID          int64  `json:"id"`
	Date        string `json:"date"`
	Provider    int64  `json:"provider"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
	CostCenter  int64  `json:"cost_center"`
	Account     int64  `json:"account"`
}

// album represents data about a record album.
type TransactionInput struct {
	ID          int64  `json:"id"`
	Date        string `json:"date"`
	Provider    int64  `json:"provider"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
	CostCenter  int64  `json:"cost_center"`
	Account     int64  `json:"account"`
}

func Flatten(transaction *Transaction) TransactionOutput {
	return TransactionOutput{
		ID:          transaction.ID,
		Date:        transaction.Date,
		Provider:    transaction.ProviderID,
		Description: transaction.Description,
		Amount:      transaction.Amount,
		CostCenter:  transaction.CostCenterID,
		Account:     transaction.AccountID,
	}

}

func (*TransactionModel) GetList(sort []string, itemRange []int, filters map[string]interface{}) []Transaction {
	var transactions []Transaction
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
	db.Find(&transactions)
	return transactions
}

func (*TransactionModel) Delete(id []*Transaction) {
	for _, i := range id {
		db.Delete(i)
	}
}

func (*TransactionModel) TotalCount() int64 {
	var count int64
	db.Model(&Transaction{}).Count(&count)
	return count
}

func (*TransactionModel) Create(newtx *Transaction) {
	db.Create(newtx)
}

func (*TransactionModel) GetSingleItem(id int64) (*Transaction, error) {
	var transaction Transaction
	if err := db.Preload(clause.Associations).Where("id = ?", id).First(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}
