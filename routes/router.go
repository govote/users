package routes

import (
	"os"

	"github.com/deputadosemfoco/users/controllers"
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
	e.Use(middleware.BodyLimit("1K"))

	healthCtrl := controllers.HealthCtrl{}
	userCtrl := buildUserController(e)

	// health check route
	e.Get("/api/chk", healthCtrl.Check)

	// user routes
	e.POST("/api/user", userCtrl.Post)

	return e
}
