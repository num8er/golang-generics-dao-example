package entities

import (
	"database/sql/driver"
	"time"
)

type IEntity interface {
}

type NullTime struct {
	time.Time
	Valid bool
}

func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

type Entity struct {
	IEntity   `json:"-"`
	Id        string   `json:"id" db:"id"`
	CreatedAt NullTime `json:"updatedAt" db:"created_at"`
	UpdatedAt NullTime `json:"createdAt,omitempty" db:"updated_at"`
	Deleted   bool     `json:"deleted,omitempty" db:"deleted"`
}
