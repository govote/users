package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureFbApiReferenceIsValid(t *testing.T) {
	FacebookAPP()
	assert.NotNil(t, FbApp)
}
