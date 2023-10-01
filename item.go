package tavern

import (
	"github.com/google/uuid"
)

type Item struct {
	// ID is identifier of the entity
	ID          uuid.UUID
	Name        string
	Description string
}
