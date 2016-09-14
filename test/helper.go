package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func CreateContext() (echo.Context, *http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	req := new(http.Request)

	req.Header = make(map[string][]string)
	rec := httptest.NewRecorder()

	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	return c, req, rec
}

func CreateJsonContext(data interface{}) (echo.Context, *httptest.ResponseRecorder, error) {
	c, _, res := CreateContext()
	b, err := json.Marshal(data)

	c.Request().SetBody(bytes.NewReader(b))
	c.Request().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)

	return c, res, err
}

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
