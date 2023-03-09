package controller

import (
	"api-golang/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Postproduct(c *gin.Context) {
	var items models.Product

	if err := c.ShouldBindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// conn, err := database.Connection()
	// shared.CheckError(err)
	// Query := `INSERT INTO "Products" (id,title,price,info,rating) VALUES ($1,$2,$3,$4,$5)`

	// for _, item := range items {
	// 	item.Id = uuid.NewString()
	// 	_, err = conn.Exec(Query, item.Id, item.Title, item.Price, item.Info, item.Rating)
	// 	shared.CheckError(err)
	// }
	jsonData, _ := json.MarshalIndent(items, "", " ")
	fmt.Println(string(jsonData))
	c.JSON(http.StatusOK, items)
}
