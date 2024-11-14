package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func initServer(e *echo.Echo) http.Handler {
	e.GET("/fortune", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return c.JSON(http.StatusOK, `{"fortune": "大吉}`)
	})
	return e
}

func main() {
	e := echo.New()
	e = initServer(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func test(t *testing.T) {
	e := echo.New()
	httptest.NewServer(e)
}
