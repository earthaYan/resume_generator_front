package repositories

import "resume_backend/models"

// struct->表结构
func migrate() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&models.User{}, &models.Resume{}, &models.ResumeBasicInfo{}, &models.ResumeCareerTarget{}, &models.ResumeEducation{}, &models.ResumeProjectExperience{}, &models.ResumeSkill{}, &models.ResumeWorkExperience{})
	if err != nil {
		panic(err)
	}
}
