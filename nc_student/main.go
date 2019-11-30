package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/NDLPhat/demo_golang/nc_student/route"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	e.GET("/", hello)
	route.All(e)
	fmt.Println(e.Start(":9090"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
