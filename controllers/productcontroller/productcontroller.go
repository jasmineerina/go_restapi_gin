package productcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasmineerina/go_restapi_gin.git/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) { // menampilkan seluruh data di product

	var products []models.Product // mengambil model bernama product

	models.DB.Find(&products)                          // mencari data dari tabel product
	c.JSON(http.StatusOK, gin.H{"products": products}) // untuk menampilkan data product
}

func Show(c *gin.Context) { // menampilkan data product by id
	var product models.Product
	id := c.Param("id") // sesuai id

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
