package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/vitorsalgado/la-democracia/auth/services"
)

// UserCtrl handles all requests related to User
type UserCtrl struct {
}

// NewUserCtrl returns a new instance of UserCtrl
func NewUserCtrl() *UserCtrl {
	return &UserCtrl{}
}

// Register registers a user if needed and returns OK if all operations were successful
func (ctrl *UserCtrl) Register(c echo.Context) error {
	req := new(services.RegistrationRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	svc := services.NewRegistrationService()
	res := svc.Register(req)

	if res.Success {
		return c.JSON(http.StatusOK, res)
	}

	return c.JSON(http.StatusBadRequest, res)
}
