package controllers

import (
	"net/http"

	"github.com/deputadosemfoco/go-libs/messages"
	"github.com/deputadosemfoco/users/interactors"
	"github.com/labstack/echo"
)

type (
	RegistrationInteractor interface {
		Register(req *interactors.RegistrationRequest) interactors.RegistrationResult
	}

	UserCtrl struct {
		Interactor RegistrationInteractor
	}
)

func (ctrl *UserCtrl) Post(c echo.Context) error {
	req := new(interactors.RegistrationRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	res := ctrl.Interactor.Register(req)

	if res.Success {
		code := http.StatusOK

		if res.Created {
			code = http.StatusCreated
		}

		return c.JSON(code, res.Data)
	}

	return c.JSON(http.StatusBadRequest, messages.Error{Message: res.Message, ValidationMessages: res.ValidationMessages})
}
