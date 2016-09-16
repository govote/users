package controllers

import (
	"net/http"
	"testing"

	"github.com/deputadosemfoco/users/test"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckMustReturnHttpStatus204NoContent(t *testing.T) {
	c, _, res := test.CreateContext()

	healthCtrl := new(HealthCtrl)
	healthCtrl.Check(c)

	assert.Equal(t, http.StatusNoContent, res.Code)
}
