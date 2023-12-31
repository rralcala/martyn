package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rralcala/martyn/models"
)

func getAccounts(c *gin.Context) {
	source := models.AccountModel{}

	count := source.TotalCount()
	sort := parseJSONArray(c.Query("sort"))
	itemRange := parseJSONArrayInt(c.Query("range"))
	filter := parseJSONMap(c.Query("filter"))

	c.Writer.Header().Set("X-Total-Count", strconv.FormatInt(count, 10))
	c.IndentedJSON(
		http.StatusOK,
		source.GetList(sort, itemRange, filter),
	)
}

func getAccountByID(c *gin.Context) {
	source := new(models.AccountModel)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id is not int"})
	}
	item, err := source.GetSingleItem(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "transaction not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, item)
}

func updateAccounts(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}

func deleteAccounts(c *gin.Context) {
	source := new(models.AccountModel)

	if len(c.Param("id")) > 0 {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "id is not int"})
			return
		}
		item := models.Account{
			ID: id,
		}
		source.Delete([]*models.Account{&item})
	} else if filter := c.Query("filter"); len(filter) > 0 {
		var items []*models.Account
		for _, id := range parseJSONArrayInt(filter) {
			items = append(items, &models.Account{
				ID: int64(id),
			})
		}
		source.Delete(items)

	}
	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}

func postAccount(c *gin.Context) {
	source := new(models.AccountModel)

	var newItem models.Account

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	source.Create(&newItem)

	c.IndentedJSON(http.StatusCreated, newItem)
}
