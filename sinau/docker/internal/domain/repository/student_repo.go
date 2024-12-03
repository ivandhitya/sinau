package repository

import (
	"context"
	"ivandhitya/docker/internal/domain/entity"
)

type StudentRepository interface {
	GetByID(ctx context.Context, id int) (*entity.Student, error)
	Save(ctx context.Context, student *entity.Student) error
}
