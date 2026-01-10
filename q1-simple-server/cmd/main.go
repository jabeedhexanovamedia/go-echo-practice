package main

// Hello World API: Create a simple Echo server with GET / that returns "Hello, World!".
import "github.com/labstack/echo/v4"

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {

		return c.String(200, "Hello world")
	})
	// e.Logger.Fatal(e.Start(":8080"))

	e.Start(":0")
}
