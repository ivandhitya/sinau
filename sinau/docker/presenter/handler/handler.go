package handler

import "github.com/labstack/echo/v4"

type StudentHandler interface {
	GetStudent(c echo.Context) error
	UpsertStudent(c echo.Context) error
}
