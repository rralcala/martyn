package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rralcala/martyn/lib/log"
	"github.com/rralcala/martyn/models"
)

func getTransactions(c *gin.Context) {
	transactions := models.GetTransacions()
	var ret []models.TransactionOutput
	for _, a := range transactions {
		ret = append(ret, models.Flatten(&a))
	}
	c.IndentedJSON(http.StatusOK, ret)
}

// getTransactionByID locates the transaction whose ID value matches the id
// parameter sent by the client, then returns that transaction as a response.
func getTransactionByID(c *gin.Context) {
	log.Info("getTransactionByID")
	// Loop over the list of transactions, looking for
	// a transaction whose ID value matches the parameter.
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id is not int"})
	}
	transaction, err := models.GetTransactionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "transaction not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, models.Flatten(transaction))
}

// postTransaction adds a transaction from JSON received in the request body.
func postTransaction(c *gin.Context) {
	var newTransaction models.TransactionInput

	// Call BindJSON to bind the received JSON to
	// newTransaction.
	if err := c.BindJSON(&newTransaction); err != nil {
		return
	}

	// Add the new album to the slice.
	transactionStruct, err := models.Build(&newTransaction)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	models.AppendTransacions(transactionStruct)

	c.IndentedJSON(http.StatusCreated, newTransaction)
}

func main() {
	router := gin.Default()

	models.ConnectDatabase()
	router.GET("/transactions", getTransactions)
	router.POST("/transactions", postTransaction)
	router.GET("/transactions/:id", getTransactionByID)
	router.Run("localhost:8080")
}
