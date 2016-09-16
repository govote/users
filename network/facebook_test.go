package network

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestEnsureFbApiReferenceIsValid(t *testing.T) {
	FacebookAPP()
	assert.NotNil(t, FbApp)
}

// #1 place a .env file in "network" folder
// #2 get a facebook access token and set env var TEST_FB_ACCESS_TOKEN in .env file
func TestValidAccessTokenShouldReturnTrue(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		return
	}

	FacebookAPP()

	accessToken := os.Getenv("TEST_FB_ACCESS_TOKEN")
	fb := FacebookGraphAPIClient{}
	res := fb.ValidateAccessToken(accessToken)

	assert.True(t, res)
}

func TestEmptyAccessTokenShouldReturnFalse(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		return
	}

	FacebookAPP()

	fb := FacebookGraphAPIClient{}
	res := fb.ValidateAccessToken("<INVALID_ACCESS_TOKEN>")

	assert.False(t, res)
}
