package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/oherych/badbadcode/pkg"
	"log"
)

func Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		t := pkg.Task{}
		err := c.Bind(&t)
		if err != nil {
			log.Fatal(err)

			return err
		}

		_, err = pkg.DB.Exec("INSERT INTO tasks(name, description)", t.TaskName, t.TaskDescription)
		if err != nil {
			return err
		}

		tm.Lock()
		tm.List[t] = ""
		tm.Unlock()

		r := new(pkg.Result)
		go tm.Go(t, r)

		return c.JSON(200, r)
	}
}
