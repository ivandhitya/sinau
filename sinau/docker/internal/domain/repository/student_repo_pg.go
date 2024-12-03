package repository

import (
	"context"
	"ivandhitya/docker/internal/domain/entity"

	"gorm.io/gorm"
)

type studentRepo struct {
	DB *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepo{
		DB: db,
	}
}

func (repo *studentRepo) GetByID(ctx context.Context, id int) (*entity.Student, error) {

	var student entity.Student
	if err := repo.DB.First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (repo *studentRepo) Save(ctx context.Context, student *entity.Student) error {
	if err := repo.DB.Save(&student).Error; err != nil {
		return err
	}
	return nil
}
