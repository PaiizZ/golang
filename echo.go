package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

func main() {
	e := echo.New()
	e.GET("/", rootHandler)
	e.Logger.Fatal(e.Start(":5002"))
}
