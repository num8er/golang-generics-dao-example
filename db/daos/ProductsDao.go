package daos

import (
	"go-generics-dao-example/db/common/mysql"
	"go-generics-dao-example/db/entities"
)

type ProductsDao struct {
	mysql.MySQLDao[entities.ProductEntity]
}

func NewProductsDao() *ProductsDao {
	instance := ProductsDao{}
	instance.SetCollectionName("products")

	return &instance
}
