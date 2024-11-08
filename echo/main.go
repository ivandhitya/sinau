package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Struct untuk representasi data User
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Dummy data User
var users = []User{
	{ID: 1, Name: "Agung", Age: 43},
	{ID: 2, Name: "Farel", Age: 37},
}

// Handler untuk mendapatkan semua user
func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

// Handler untuk mendapatkan user berdasarkan ID
func getUser(c echo.Context) error {
	id := c.Param("id")
	idTemp, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid parameter")
	}

	for _, user := range users {
		if user.ID == idTemp {
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.JSON(http.StatusNotFound, "User not found")
}

// Handler untuk membuat atau menambahkan user baru
func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	users = append(users, *user)
	return c.JSON(http.StatusCreated, user)
}

func main() {
	e := echo.New()

	// Routing
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.POST("/users", createUser)

	// Start server
	e.Logger.Fatal(e.Start(":7777"))
}
