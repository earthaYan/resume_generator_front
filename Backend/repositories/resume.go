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
	err = s.DB.Preload("BaseInfo").
		Preload("CareerTarget").
		Preload("EducationInfo").
		Preload("ProjectExperience").
		Preload("SkillInfo").
		Preload("WorkExperience").
		First(&r, "id = ? AND user_id = ?", ResumeId, uid).Error
	return
}

func (s *ResumeDao) UpdateResume(uid int64, req *models.UpdateResumeReq) error {
	// 判断该简历是否存在
	err := s.DB.Model(&models.Resume{}).Where("id=? AND uid=?", req.Id, uid).Updates(req.Updated).Error
	return err
}

func (s *ResumeDao) DeleteResume(uid int64, req *models.DeleteResumeReq) error {
	var resume models.Resume
	err := s.DB.Preload("BaseInfo").
		Preload("CareerTarget").
		Preload("EducationInfo").
		Preload("ProjectExperience").
		Preload("SkillInfo").
		Preload("WorkExperience").First(&resume, "id = ? AND user_id = ?", req.Id, uid).Error
	if err != nil {
		return err
	}
	// 采用大事物，可以回滚
	tx := s.DB.Begin()
	if err := tx.Delete(&resume).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 删除关联表数据
	if err := tx.Unscoped().Delete(&resume.BaseInfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Delete(&resume.CareerTarget).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Delete(&resume.EducationInfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Delete(&resume.SkillInfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Delete(&resume.ProjectExperience).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Delete(&resume.WorkExperience).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
