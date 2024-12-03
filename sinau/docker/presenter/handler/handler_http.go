package handler

import (
	"context"
	"net/http"
	"strconv"

	"ivandhitya/docker/internal/usecase"
	"ivandhitya/docker/presenter/model"

	"github.com/labstack/echo/v4"
)

type studentHandler struct {
	useCase usecase.StudentUseCase
}

func NewStudentHandler(uc usecase.StudentUseCase) StudentHandler {
	return &studentHandler{useCase: uc}
}

// @Summary		Get student by ID
// @Description	Retrieve student data by ID, either from cache or database
// @Tags			student
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Student ID"
// @Success		200	{object}	model.Student
// @Failure		400	{string}	string	"Missing 'id' parameter"
// @Failure		404	{string}	string	"Student not found"
// @Failure		500	{string}	string	"Error accessing cache"
// @Router			/student/{id} [get]
func (h *studentHandler) GetStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	student, err := h.useCase.GetStudent(context.Background(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Student not found")
	}
	return c.JSON(http.StatusOK, student)
}

// UpsertStudent updates a student's information.
//
//	@Summary		Upsert a student
//	@Description	Upsert student data in the database
//	@Tags			student
//	@Accept			json
//	@Produce		json
//	@Param			student	body		model.Student	true	"Student data"
//	@Success		200		{string}	string			"Student updated"
//	@Failure		400		{string}	string			"Invalid request"
//	@Failure		500		{string}	string			"Failed to Upsert student"
//	@Router			/student [put]
func (h *studentHandler) UpsertStudent(c echo.Context) error {
	var student model.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	err := h.useCase.UpsertStudent(context.Background(), &student)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to Upsert student")
	}
	return c.JSON(http.StatusOK, "Student updated")
}
