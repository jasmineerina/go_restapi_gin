package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jasmineerina/go_restapi_gin.git/controllers/productcontroller"
	"github.com/jasmineerina/go_restapi_gin.git/models"
)

func main() {
	r := gin.Default()         // inisialisai router gin
	models.ConnectedDatabase() // untuk connect ke database

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products", productcontroller.Update)
	r.DELETE("/api/products", productcontroller.Delete)

	r.Run() // menjalankan server

}
