package main

import (
	"github.com/labstack/echo/v4"
	"github.com/oherych/badbadcode/handlers"
	"github.com/oherych/badbadcode/pkg"
)

func main() {
	tm := new(pkg.TaskManager)

	e := echo.New()
	e.POST("/task", handlers.Create(tm))

	if err := e.Start(":1323"); err != nil {
		panic(err)
	}
}
