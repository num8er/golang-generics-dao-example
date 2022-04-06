package customers

import (
	"github.com/gin-gonic/gin"

	"go-generics-dao-example/db/daos"
)

func Get(c *gin.Context) {
	id := c.Param("id")

	dao := daos.NewCustomersDao()
	customer, err := dao.FindById(id)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	if customer == nil {
		c.JSON(404, gin.H{})
		return
	}

	c.JSON(200, customer)
}
