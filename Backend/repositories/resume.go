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
func (s *ResumeDao) ListResume(start, limit int, uid int64) (r []*models.Resume, total int64, err error) {
	err = s.DB.Model(&models.Resume{}).Where("user_id=?", uid).Count(&total).Error
	if err != nil {
		return
	}
	err = s.DB.Model(&models.Resume{}).Where("user_id=?", uid).Limit(limit).Offset((start - 1) * limit).Find(&r).Error
	return

}
func (s *ResumeDao) FindResumeByResumeIdAndUserId(ResumeId string, uid int64) (r *models.Resume, err error) {
	// err = s.DB.Model(&models.Resume{}).Where("user_id=? AND id=?", uid, ResumeId).First(&r).Error
	// return
	err = s.DB.Preload("BaseInfo").
		Preload("CareerTarget").
		Preload("EducationInfo").
		Preload("ProjectExperience").
		Preload("SkillInfo").
		Preload("WorkExperience").
		First(&r, "id = ? AND user_id = ?", ResumeId, uid).Error
	return
}
