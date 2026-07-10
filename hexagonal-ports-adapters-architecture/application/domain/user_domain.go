package domain

import "time"

type User struct {
	ID        int32     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Age       int32     `json:"age"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
