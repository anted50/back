package controller

import (
	"math"
	"strconv"
	"test/initializers"
	"test/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Get(c *gin.Context) {
	var req models.Get
	var response models.Trip
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	db := initializers.DBGorm
	var trips []models.Trip
	db.Where("id = ?", req.ID).Find(&trips)
	for _, t := range trips {
		if t.ID == req.ID {
			response.ID = t.ID
			response.Title = t.Title
			response.Desc = t.Desc
			response.Url = t.Url
			response.Price = t.Price
			response.Capacity = t.Capacity
			response.Type = t.Type
			response.HasCompFood = t.HasCompFood
			response.Additional = t.Additional
			response.ContactNo = t.ContactNo
		}
	}
	c.JSON(200, response)
}

func Paginate(c *gin.Context) {
	//TODO : Give how many pages there would be
	pageStr := c.Param("id")
	page, _ := strconv.Atoi(pageStr)
	offset := (page - 1) * 10

	var results []models.Trip
	var count int64

	db := initializers.DBGorm

	db.Model(&results).Count(&count)
	db.Limit(5).Offset(offset).Find(&results)

	TotalPage := math.Ceil(float64(count) / 5)
	lastPage := 0
	if page < 1 {
		lastPage = 1
	} else {
		lastPage = page - 1
	}
	c.JSON(200, gin.H{
		"results": results,
		"pagination": models.PaginationData{
			NextPage:    page + 1,
			LastPage:    lastPage,
			CurrentPage: page,
			TotalPage:   int(TotalPage),
		},
	})
}
