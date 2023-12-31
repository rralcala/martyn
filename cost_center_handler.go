package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rralcala/martyn/models"
)

func getCostCenters(c *gin.Context) {
	source := models.CostCenterModel{}

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

func getCostCenterByID(c *gin.Context) {
	source := new(models.CostCenterModel)

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

func updateCostCenters(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}

func deleteCostCenters(c *gin.Context) {
	source := new(models.CostCenterModel)

	if len(c.Param("id")) > 0 {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "id is not int"})
			return
		}
		item := models.CostCenter{
			ID: id,
		}
		source.Delete([]*models.CostCenter{&item})
	} else if filter := c.Query("filter"); len(filter) > 0 {
		var items []*models.CostCenter
		for _, id := range parseJSONArrayInt(filter) {
			items = append(items, &models.CostCenter{
				ID: int64(id),
			})
		}
		source.Delete(items)

	}
	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}

func postCostCenter(c *gin.Context) {
	source := new(models.CostCenterModel)

	var newItem models.CostCenter

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	source.Create(&newItem)

	c.IndentedJSON(http.StatusCreated, newItem)
}
