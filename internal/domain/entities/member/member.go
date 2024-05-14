package member

import (
	"time"

	"github.com/google/uuid"
)

// Member Model that represents the Member
type Member struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}
