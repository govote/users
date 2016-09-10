package services

import "github.com/vitorsalgado/la-democracia/lib/go/messages"

// RegistrationRequest represents the required information
// to register of login a User
type RegistrationRequest struct {
	messages.Request
	Name       string
	Email      string
	PhotoURL   string
	FacebookID string
}
