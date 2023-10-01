package tavern

import (
	"github.com/google/uuid"
)

type Person struct {
	// ID is identifier of the entity
	ID   uuid.UUID
	Name string
	Age  int
}
