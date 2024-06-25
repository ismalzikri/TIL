package main

import (
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Gzip())

	e.GET("/image", func(c echo.Context) error {
		f, err := os.Open("tes.jpeg")
		if err != nil {
			return err
		}

		_, err = io.Copy(c.Response(), f)
		if err != nil {
			return err
		}

		return nil
	})

	e.Logger.Fatal(e.Start(":9000"))
}
