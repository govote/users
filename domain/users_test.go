package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureUserWasConstructedCorrectly(t *testing.T) {
	user := NewUser("email@email.com", "test name", "fbid")

	assert.NotEmpty(t, user.Email)
	assert.NotEmpty(t, user.Name)
	assert.NotEmpty(t, user.FacebookID)
}
