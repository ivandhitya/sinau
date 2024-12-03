package usecase

import (
	"context"
	"ivandhitya/docker/internal/domain/entity"
	"ivandhitya/docker/presenter/model"
)

type StudentUseCase interface {
	GetStudent(ctx context.Context, id int) (*entity.Student, error)
	UpsertStudent(ctx context.Context, student *model.Student) error
}
