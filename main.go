package main

import (
	"github.com/labstack/echo/v4"
	"github.com/oherych/badbadcode/handlers"
)

func main() {
	e := echo.New()
	e.GET("/task", handlers.List())
	e.POST("/task", handlers.Create())

	if err := e.Start(":1323"); err != nil {
		panic(err)
	}
}
