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
	"github.com/labstack/echo/middleware"
)

// SetUp all application routes
func SetUp() *echo.Echo {
	e := echo.New()

	e.SetDebug(os.Getenv("GO_ENV") == "development")

	e.Use(middleware.Recover())
	e.SetHTTPErrorHandler(errorHandler)
	e.Use(middleware.Gzip())

	healthCtrl := controllers.HealthCtrl{}
	userCtrl := buildUserController()

	// health check route
	e.Get("/api/chk", healthCtrl.Check)

	// user routes
	e.POST("/api/user", userCtrl.Post)

	return e
}

func buildUserController() *controllers.UserCtrl {
	registrationInteractor := new(interactors.RegistrationInteractor)
	registrationInteractor.FacebookGraphAPI = new(network.FacebookGraphAPIClient)
	registrationInteractor.UserRepository = new(repositories.SQLUserRepository)

	userCtrl := new(controllers.UserCtrl)
	userCtrl.Interactor = registrationInteractor

	return userCtrl
}

func errorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := err.Error()
	detail := ""

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}

	if msg == "" {
		msg = http.StatusText(code)
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
