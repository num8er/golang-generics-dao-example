package mysql

import (
	"context"
	"strconv"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/georgysavva/scany/sqlscan"

	"go-generics-dao-example/db/common/daos"
	"go-generics-dao-example/db/entities"
)

type MySQLDao[T entities.IEntity] struct {
	daos.Dao[T]
	collectionName string
}

func (d *MySQLDao[T]) SetCollectionName(name string) {
	d.collectionName = name
}

func (d *MySQLDao[T]) CollectionName() string {
	return d.collectionName
}

func (d *MySQLDao[T]) FindById(id string) (*T, error) {
	var dest []*T

	query, _, _ := goqu.Dialect("mysql").From(d.CollectionName()).Where(goqu.Ex{"id": id}).Limit(1).ToSQL()
	err := sqlscan.Select(context.Background(), DB, &dest, query)
	if err != nil {
		return nil, err
	}

	if len(dest) == 1 {
		return dest[0], nil
	}
	return nil, nil
}

func (d *MySQLDao[T]) Create(data interface{}) (*T, error) {
	query, _, _ := goqu.Dialect("mysql").Insert(d.CollectionName()).Rows(data).ToSQL()
	result, err := DB.Exec(query)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	return d.FindById(strconv.FormatInt(id, 10))
}
