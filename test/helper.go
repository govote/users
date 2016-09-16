package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

// CreateContext creates a fake Echo HttpContext for tests
func CreateContext() (echo.Context, *http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	req := new(http.Request)

	req.Header = make(map[string][]string)
	rec := httptest.NewRecorder()

	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	return c, req, rec
}

// CreateJSONContext creates a fake Echo HttpContext that accepts and responds json
func CreateJSONContext(data interface{}) (echo.Context, *httptest.ResponseRecorder, error) {
	c, _, res := CreateContext()
	b, err := json.Marshal(data)

	c.Request().SetBody(bytes.NewReader(b))
	c.Request().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)

	return c, res, err
}

// PanicErr panics if err is not nil
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
