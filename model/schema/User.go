package schema

import (
	"github.com/google/uuid"
)

type User struct {
	Id    uuid.UUID `from:"Id"`
	Name  string    `json:"Name" binding:"required"`
	Email string    `json:"Email" binding:"required"`
	Phone string    `json:"Phone" binding:"required"`
}
