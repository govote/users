package controllers

import (
	"net/http"

	"github.com/deputadosemfoco/go-libs/messages"
	"github.com/deputadosemfoco/users/interactors"
	"github.com/labstack/echo"
)

// AuthCtrl ...
type AuthCtrl struct {
	Interactor interactors.AuthInteractorContract
}

func (ctrl *AuthCtrl) Post(c echo.Context) error {
	req := new(interactors.AuthRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	res, err := ctrl.Interactor.Authenticate(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, messages.Error{Message: err.Error()})
	}

	if res.Authenticated {
		return c.JSON(http.StatusOK, res.User)
	}

	return c.JSON(http.StatusUnauthorized, nil)
}
