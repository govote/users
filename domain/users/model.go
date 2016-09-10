package users

import (
	"time"

	"github.com/satori/go.uuid"
	"github.com/vitorsalgado/la-democracia/lib/go/domain"
)

type (
	// User represents the application user
	User struct {
		domain.Entity

		ID         string `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		PhotoURL   string `json:"photoUrl"`
		FacebookID string `json:"facebookId"`
		GoogleID   string `json:"googleId"`
	}

	// OptionalUser is a container which may or may not contain a User
	OptionalUser struct {
		Valid bool
		User  User
	}
)

// NewUser is the preferred way to build a User struct
func NewUser(name, email string) User {
	return User{
		ID:    uuid.NewV4().String(),
		Name:  name,
		Email: email,

		Entity: domain.Entity{CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()},
	}
}
