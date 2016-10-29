package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/deputadosemfoco/go-libs/messages"
	"github.com/deputadosemfoco/go-libs/test"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestEnsureRouteSetUp(t *testing.T) {
	e := SetUp()

	assert.NotNil(t, e)
	assert.False(t, e.Debug())

	routes := e.Routes()

	for _, r := range routes {
		assert.True(t, strings.HasPrefix(r.Path, "/api"))
	}
}

func TestEnsureUserControllerIsOK(t *testing.T) {
	userCtrl := buildUserController()

	assert.NotNil(t, userCtrl)
	assert.NotNil(t, userCtrl.Interactor)
}

func TestRegularError(t *testing.T) {
	err := errors.New("regular error")
	c, _, res := test.CreateContext()
	model := new(messages.Error)

	errorHandler(err, c)
	json.Unmarshal(res.Body.Bytes(), model)

	assert.NotNil(t, model)
	assert.Equal(t, "regular error", model.Message)
	assert.Equal(t, http.StatusInternalServerError, model.Code)
	assert.Equal(t, http.StatusInternalServerError, res.Code)
}

func TestEchoError(t *testing.T) {
	c, _, res := test.CreateContext()
	model := new(messages.Error)

	errorHandler(&echo.HTTPError{Code: 10, Message: "echo error"}, c)
	json.Unmarshal(res.Body.Bytes(), model)

	assert.NotNil(t, model)
	assert.Equal(t, "echo error", model.Message)
	assert.Equal(t, 10, model.Code)
	assert.Equal(t, 10, res.Code)
}
