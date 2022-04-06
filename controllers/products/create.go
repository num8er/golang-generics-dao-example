package products

import (
	"github.com/gin-gonic/gin"

	"go-generics-dao-example/db/daos"
	"go-generics-dao-example/db/entities"
)

func Create(c *gin.Context) {
	product := entities.Product{}
	_ = c.BindJSON(&product)

	dao := daos.NewProductsDao()
	newProduct, err := dao.Create(&product)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(200, newProduct)
}
