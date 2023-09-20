package models

import "gorm.io/gorm"

type Resume struct {
	gorm.Model
	UserId            uint                      `json:"user_id"`
	BaseInfo          ResumeBasicInfo           `json:"base_info"`
	CareerTarget      ResumeCareerTarget        `json:"career_target"`
	EducationInfo     []ResumeEducation         `json:"education"`
	ProjectExperience []ResumeProjectExperience `json:"project_experience"`
	SkillInfo         []ResumeSkill             `json:"skills"`
	WorkExperience    []ResumeWorkExperience    `json:"work_experience"`
}
type ResumeBasicInfo struct {
	gorm.Model
	ResumeId        uint   `json:"resume_id" gorm:"not null"`
	Name            string `json:"name" gorm:"not null"`
	Age             int    `json:"age" gorm:"not null"`
	ExperienceYears int    `json:"experience_years"`
	Email           string `json:"email" gorm:"not null"`
	Phone           string `json:"phone" gorm:"not null"`
	BriefInfo       string `json:"brief_introduction"`
}
type ResumeCareerTarget struct {
	gorm.Model
	ResumeId uint   `json:"resume_id" gorm:"not null"`
	Position string `json:"position"`
	City     string `json:"city"`
	Salary   string `json:"salary"`
	Status   string `json:"status"`
}
type ResumeEducation struct {
	gorm.Model
	ResumeId uint   `json:"resume_id" gorm:"not null"`
	Time     string `json:"time"`
	School   string `json:"school"`
	Major    string `json:"major"`
	Degree   string `json:"degree"`
}
type ResumeProjectExperience struct {
	gorm.Model
	ResumeId    uint   `json:"resume_id" gorm:"not null"`
	Time        string `json:"time"`
	ProjectInfo string `json:"project_info"`
}
type ResumeSkill struct {
	gorm.Model
	ResumeId uint   `json:"resume_id" gorm:"not null"`
	Name     string `json:"skill_name"`
	Percent  string `json:"percent"`
}
type ResumeWorkExperience struct {
	gorm.Model
	ResumeId uint   `json:"resume_id" gorm:"not null"`
	Time     string `json:"time"`
	Position string `json:"position"`
	Company  string `json:"company"`
}
type CreateResumeReq struct {
	UserId            uint                      `json:"user_id" binding:"required"`
	BaseInfo          ResumeBasicInfo           `json:"base_info"`
	CareerTarget      ResumeCareerTarget        `json:"career_target"`
	EducationInfo     []ResumeEducation         `json:"education"`
	ProjectExperience []ResumeProjectExperience `json:"project_experience"`
	SkillInfo         []ResumeSkill             `json:"skills"`
	WorkExperience    []ResumeWorkExperience    `json:"work_experience"`
}
