package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vitorsalgado/la-democracia/auth/controllers"
)

// SetUp all application routes
func SetUp() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Gzip())

	userCtrl := controllers.NewUserCtrl()
	healthCtrl := controllers.NewHealthCtrl()

	e.SetHTTPErrorHandler(controllers.GtHTTPErrorHandler)

	e.POST("/api/user", userCtrl.Register)

	e.Get("/api/chk", healthCtrl.Check)

	return e
}
