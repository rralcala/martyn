package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rralcala/martyn/models"
)

func getTransactions(c *gin.Context) {
	source := models.TransactionModel{}

	count := source.TotalCount()
	sort := parseJSONArray(c.Query("sort"))
	itemRange := parseJSONArrayInt(c.Query("range"))
	filter := parseJSONMap(c.Query("filter"))
	transactions := source.GetList(sort, itemRange, filter)
	var ret []models.TransactionOutput
	for _, a := range transactions {
		ret = append(ret, models.Flatten(&a))
	}
	c.Writer.Header().Set("X-Total-Count", strconv.FormatInt(count, 10))
	c.IndentedJSON(http.StatusOK, ret)
}

// getTransactionByID locates the transaction whose ID value matches the id
// parameter sent by the client, then returns that transaction as a response.
func getTransactionByID(c *gin.Context) {
	source := new(models.TransactionModel)

	// Loop over the list of transactions, looking for
	// a transaction whose ID value matches the parameter.
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id is not int"})
	}
	transaction, err := source.GetSingleItem(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "transaction not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, models.Flatten(transaction))
}

func updateTransaction(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}

func deleteTransactions(c *gin.Context) {
	source := new(models.TransactionModel)

	if len(c.Param("id")) > 0 {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "id is not int"})
			return
		}
		item := models.Transaction{
			ID: id,
		}
		source.Delete([]*models.Transaction{&item})
	} else if filter := c.Query("filter"); len(filter) > 0 {
		var items []*models.Transaction
		for _, id := range parseJSONArrayInt(filter) {
			items = append(items, &models.Transaction{
				ID: int64(id),
			})
		}
		source.Delete(items)

	}
	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}

// postTransaction adds a transaction from JSON received in the request body.
func postTransaction(c *gin.Context) {
	source := new(models.TransactionModel)

	var newTransaction models.Transaction

	// Call BindJSON to bind the received JSON to
	// newTransaction.
	if err := c.BindJSON(&newTransaction); err != nil {
		return
	}

	source.Create(&newTransaction)

	c.IndentedJSON(http.StatusCreated, newTransaction)
}
