package entities

type Customer struct {
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
}

type CustomerEntity struct {
	Entity
	Customer
}
