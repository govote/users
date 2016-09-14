package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureUserWasConstructedCorrectly(t *testing.T) {
	user := NewUser("email@email.com", "test name")

	assert.NotEmpty(t, user.Email)
	assert.NotEmpty(t, user.Name)
}
