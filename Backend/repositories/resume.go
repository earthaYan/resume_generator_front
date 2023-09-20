package repositories

import (
	"context"
	"resume_backend/models"

	"gorm.io/gorm"
)

type ResumeDao struct {
	*gorm.DB
}

func NewResumeDao(ctx context.Context) *ResumeDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &ResumeDao{NewDBClient(ctx)}
}

func (s *ResumeDao) CreateResume(resume *models.Resume) error {
	return s.Model(&models.Resume{}).Create(resume).Error
}
