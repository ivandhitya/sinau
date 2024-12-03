package usecase

import (
	"context"
	"ivandhitya/docker/internal/domain/entity"
	"ivandhitya/docker/internal/domain/repository"
	"ivandhitya/docker/presenter/model"
)

type studentUseCase struct {
	repo repository.StudentRepository
}

func NewStudentUseCase(repo repository.StudentRepository) StudentUseCase {
	return &studentUseCase{repo: repo}
}

func (u *studentUseCase) GetStudent(ctx context.Context, id int) (*entity.Student, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *studentUseCase) UpsertStudent(ctx context.Context, student *model.Student) error {
	tempStudent := &entity.Student{
		ID:    student.ID,
		Name:  student.Name,
		Grade: student.Grade,
	}
	return u.repo.Save(ctx, tempStudent)
}
