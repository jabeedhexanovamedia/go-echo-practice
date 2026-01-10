package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type UserSignupRequest struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email"`
}

// Simple JSON Response: Create GET /user that returns a hardcoded user JSON: { "id":1, "name":"John" }.
func main() {
	e := echo.New()

	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"id":   1,
			"name": "john",
		})
	})

	// using echo.Map
	e.GET("/user", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"id":   1,
			"name": "John",
		})
	})

	// using struct
	e.GET("/users2", func(C echo.Context) error {
		return C.JSON(200, User{
			Id:   2,
			Name: "john 2",
		})
	})

	// 	Create an Echo API `POST /users` that:
	// - Accepts JSON request body
	// - Converts JSON → Go struct
	// - Returns the same data as JSON response

	e.POST("/users", func(c echo.Context) error {

		var userReq UserSignupRequest

		// JSON → Go struct
		if err := c.Bind(&userReq); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "invalid request payload",
			})
		}

		// Go struct → JSON
		return c.JSON(200, userReq)

	})
	e.Logger.Fatal(e.Start("127.0.0.1:3000"))
}
