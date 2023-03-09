package main

import (
	"api-golang/controller"
	"log"

	// "api-golang/migrator"

	"github.com/gin-gonic/gin"
)

func main() {
	// db, _ := database.Connection()
	// migrator.Migrate(db)
	// defer db.Close()

	router := gin.Default()
	router.POST("/products", controller.Postproduct)
	log.Fatal(router.Run(":8000"))
}
