package handler

import (
	"ivandhitya/docker/internal/domain/repository"
	"ivandhitya/docker/internal/usecase"
	"ivandhitya/docker/pkg/database"

	_ "ivandhitya/docker/presenter/handler/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RoutingRestAPI(e *echo.Echo) error {

	// Dependency Injection
	dbConn, err := database.NewPostgresConnection("postgres", "sinau", "12345", "app_db", "5432")
	if err != nil {
		e.Logger.Error(err)
		return err
	}
	studentRepo := repository.NewStudentRepository(dbConn)
	studentUC := usecase.NewStudentUseCase(studentRepo)
	studentHandler := NewStudentHandler(studentUC)

	// Routing
	e.GET("/student/:id", studentHandler.GetStudent)
	e.PUT("/student", studentHandler.UpsertStudent)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return nil
}
