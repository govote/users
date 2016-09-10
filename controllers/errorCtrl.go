package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/vitorsalgado/la-democracia/lib/go/messages"
)

// GtHTTPErrorHandler is a centralized error handler for whole application
func GtHTTPErrorHandler(err error, c echo.Context) {
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
