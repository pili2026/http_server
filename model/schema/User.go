package schema

import (
	"github.com/google/uuid"
)

type User struct {
	Id    uuid.UUID `from:"Id"`
	Name  string    `json:"Name"`
	Email string    `json:"Email"`
	Phone string    `json:"Phone"`
}
