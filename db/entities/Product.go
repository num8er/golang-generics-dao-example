package entities

type Product struct {
	Name string `json:"name" db:"name"`
}

type ProductEntity struct {
	Entity
	Product
}
