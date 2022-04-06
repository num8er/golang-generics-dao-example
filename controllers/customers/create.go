package customers

import (
	"github.com/gin-gonic/gin"

	"go-generics-dao-example/db/daos"
	"go-generics-dao-example/db/entities"
)

func Create(c *gin.Context) {
	customer := entities.Customer{}
	_ = c.BindJSON(&customer)

	dao := daos.NewCustomersDao()
	newCustomer, err := dao.Create(&customer)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(200, newCustomer)
}
