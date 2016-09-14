package routes

import (
	"net/http"

	"github.com/deputadosemfoco/go-libs/messages"
	"github.com/deputadosemfoco/users/controllers"
	"github.com/deputadosemfoco/users/interactors"
	"github.com/deputadosemfoco/users/repositories"
	"github.com/labstack/echo"
)

func BuildUserController(e *echo.Echo) *controllers.UserCtrl {
	registrationInteractor := new(interactors.RegistrationInteractor)
	registrationInteractor.UserRepository = new(repositories.SQLUserRepository)

	userCtrl := new(controllers.UserCtrl)
	userCtrl.Interactor = registrationInteractor

	return userCtrl
}

func ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := http.StatusText(code)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}

	if msg == "" {
		msg = err.Error()
	}

	resp := messages.Response{Success: false, Message: msg}

	if !c.Response().Committed() {
		c.JSON(code, resp)
	}
}
