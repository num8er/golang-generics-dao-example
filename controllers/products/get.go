package products

import (
	"github.com/gin-gonic/gin"

	"go-generics-dao-example/db/daos"
)

func Get(c *gin.Context) {
	id := c.Param("id")

	dao := daos.NewProductsDao()
	product, err := dao.FindById(id)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	if product == nil {
		c.JSON(404, gin.H{})
		return
	}

	c.JSON(200, product)
}
