package routes

import (
	"github.com/deputadosemfoco/users/controllers"
	"github.com/labstack/echo"
)

// SetUp all application routes
func SetUp() *echo.Echo {
	e := echo.New()
	e.SetHTTPErrorHandler(ErrorHandler)

	healthCtrl := controllers.HealthCtrl{}
	userCtrl := BuildUserController(e)

	// health check route
	e.Get("/api/chk", healthCtrl.Check)

	// user routes
	e.POST("/api/user", userCtrl.Post)

	return e
}
