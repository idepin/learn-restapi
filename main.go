package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	loadEnv()
	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:id", getUser)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}

	e.GET("/pass", getPass)

}

func getPass(c *echo.Context) error {
	pass := os.Getenv("PASS")
	return c.String(http.StatusOK, pass)
}

func getUser(c *echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
