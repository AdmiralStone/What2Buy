package handler

import (
	"What2Buy/Server/helper"
	"What2Buy/Server/models"
	"time"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetItems(c *gin.Context) {
	var itemList []models.Item
	// Instantiate the db connection variable
	db := helper.Helper()

	// if err != nil {
	// 	panic(err.Error())
	// }

	results, err := db.Query("SELECT itemId,itemName,itemVotes FROM product_dtl")
	if err != nil {
		panic(err.Error())
	}

	// Iterate over quesry result and append each row to itemList
	for results.Next() {
		var item models.Item
		err = results.Scan(&item.ItemID, &item.ItemName, &item.ItemVotes)
		if err != nil {
			panic(err.Error())
		}
		itemList = append(itemList, item)
	}
	defer db.Close()
	defer results.Close()

	// Return Item List to User
	c.JSON(200, &itemList)

}

func AddItem(c *gin.Context) {
	var item models.Item
	c.Bind(&item)

	db := helper.Helper()

	insert, err := db.Query("INSERT INTO product_dtl (itemName,itemVotes,createdOn) VALUES (?,?,?)", item.ItemName, item.ItemVotes, time.Now())

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	defer db.Close()

	GetItems(c)

}

func UpvoteItem(c *gin.Context) {
	var item models.Item
	c.Bind(&item)

	db := helper.Helper()

	res, err := db.Query("UPDATE product_dtl SET itemVotes = itemVotes+1 WHERE itemId = ?", item.ItemID)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	defer res.Close()

	GetItems(c)

}

func DeleteItem(c *gin.Context) {
	var item models.Item
	c.Bind(&item)

	db := helper.Helper()
	_, err := db.Query("DELETE FROM product_dtl WHERE itemId = ?", item.ItemID)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	GetItems(c)
}
