package models

import (
	"errors"
	"strings"

	"github.com/rralcala/martyn/lib/log"
	"gorm.io/gorm/clause"
)

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
	Provider    string `json:"provider"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
	CostCenter  string `json:"cost_center"`
	Account     string `json:"account"`
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
		Provider:    transaction.Provider.Name,
		Description: transaction.Description,
		Amount:      transaction.Amount,
		CostCenter:  transaction.CostCenter.Description,
		Account:     transaction.Account.Description,
	}

}

func Build(transaction *TransactionInput) (*Transaction, error) {
	provider := FindProvider(transaction.Provider)
	if provider == nil {
		return nil, errors.New("Provider not found")
	}
	costCenter := FindCostCenter(transaction.CostCenter)
	if costCenter == nil {
		return nil, errors.New("costCenter not found")
	}
	account := FindAccount(transaction.Account)
	if account == nil {
		return nil, errors.New("account not found")
	}
	return &Transaction{
		ID:           transaction.ID,
		Date:         transaction.Date,
		Provider:     provider,
		ProviderID:   transaction.Provider,
		Description:  transaction.Description,
		Amount:       transaction.Amount,
		CostCenterID: transaction.CostCenter,
		CostCenter:   costCenter,
		Account:      account,
		AccountID:    transaction.Account,
	}, nil

}

func GetTransacions(sort []string, itemRange []int, filters map[string]interface{}) []Transaction {
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

func DeleteTransactions(id []*Transaction) {
	for _, i := range id {
		db.Delete(i)
	}
}

func AppendTransacions(newtx *Transaction) {
	transaction := Transaction{
		Date:         newtx.Date,
		ProviderID:   newtx.ProviderID,
		Provider:     nil,
		Description:  newtx.Description,
		Amount:       newtx.Amount,
		CostCenterID: newtx.CostCenterID,
		CostCenter:   nil,
		AccountID:    newtx.AccountID,
		Account:      nil,
	}
	db.Create(&transaction)

}

func GetTransactionByID(id int64) (*Transaction, error) {
	var transaction Transaction
	if err := db.Preload(clause.Associations).Where("id = ?", id).First(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}
