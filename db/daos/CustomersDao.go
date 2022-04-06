package daos

import (
	"go-generics-dao-example/db/common/mysql"
	"go-generics-dao-example/db/entities"
)

type CustomersDao struct {
	mysql.MySQLDao[entities.CustomerEntity]
}

func NewCustomersDao() *CustomersDao {
	instance := CustomersDao{}
	instance.SetCollectionName("customers")

	return &instance
}
