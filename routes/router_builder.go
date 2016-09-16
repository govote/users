package routes

import (
	"net/http"
	"os"
	"runtime"

	"github.com/deputadosemfoco/go-libs/messages"
	"github.com/deputadosemfoco/users/controllers"
	"github.com/deputadosemfoco/users/interactors"
	"github.com/deputadosemfoco/users/network"
	"github.com/deputadosemfoco/users/repositories"
	"github.com/labstack/echo"
)

func buildUserController(e *echo.Echo) *controllers.UserCtrl {
	registrationInteractor := new(interactors.RegistrationInteractor)
	registrationInteractor.FacebookGraphAPI = new(network.FacebookGraphAPIClient)
	registrationInteractor.UserRepository = new(repositories.SQLUserRepository)

	userCtrl := new(controllers.UserCtrl)
	userCtrl.Interactor = registrationInteractor

	return userCtrl
}

func errorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := http.StatusText(code)
	detail := ""

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}

	if msg == "" {
		msg = err.Error()
	}

	if os.Getenv("GO_ENV") == "development" {
		b := make([]byte, 2048)
		n := runtime.Stack(b, false)

		detail = string(b[:n])
	}

	resp := messages.Error{Message: msg, Code: code, Detail: detail}

	if !c.Response().Committed() {
		c.JSON(code, resp)
	}
}
