package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/oherych/badbadcode/pkg"
)

func List() echo.HandlerFunc {
	return func(c echo.Context) error {
		r := []pkg.Task{}
		for item, _ := range tm.List {
			r = append(r, item)
		}

		return c.JSON(200, r)
	}
}
