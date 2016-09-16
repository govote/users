package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// HealthCtrl is the controller responsible for health verifications
type HealthCtrl struct{}

// Check simple returns an OK if application is running
func (ctrl *HealthCtrl) Check(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
