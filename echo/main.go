package main

import (
	"net/http"
	"strconv"

	_ "github.com/ivandhitya/sinau/echo/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
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

// @Summary Mendapatkan daftar pengguna
// @Description Mendapatkan seluruh pengguna
// @Tags         users
// @Success 200 {object} map[string]interface{}
// @Router /users [get]
func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

// @Summary      Get user by ID
// @Description  Get details for a specific user by their ID
// @Tags         users
// @Param        id   path      int     true  "User ID"
// @Success      200  {object}  User
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Router       /users/{id} [get]
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

// createUser godoc
// @Summary      Create a new user
// @Description  This endpoint allows creating a new user with the provided details
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      User  true  "User Data"
// @Success      201   {object}  User
// @Failure      400   {object}  string  "Bad Request"
// @Router       /users [post]
func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	users = append(users, *user)
	return c.JSON(http.StatusCreated, user)
}

// @title Sample API
// @version 1.0
// @description Ini adalah dcontoh dokumentasi API
func main() {
	e := echo.New()

	// Routing
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.POST("/users", createUser)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hallo aku server echo")
	})

	// Start server
	e.Logger.Fatal(e.Start(":7777"))
}
