package daos

import (
	"go-generics-dao-example/db/entities"
)

type Dao[T entities.IEntity] interface {
	SetCollectionName(name string)
	CollectionName() string
	FindById(id string) (*T, error)
}
