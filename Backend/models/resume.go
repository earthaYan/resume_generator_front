package models

type ResumeBasicInfo struct {
	ID              int    `json:"id"`
	ResumeId        int    `json:"resumeId"`
	Name            string `json:"name"`
	Age             int    `json:"age"`
	ExperienceYears int    `json:"experienceYears"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	BriefInfo       string `json:"briefIntroduction"`
}
type ResumeCareerObjective struct {
	ID       int    `json:"id"`
	ResumeId int    `json:"resumeId"`
	Position string `json:"position"`
	City     string `json:"city"`
	Salary   string `json:"salary"`
	Status   string `json:"status"`
}
type ResumeEducation struct {
	ID       int    `json:"id"`
	ResumeId int    `json:"resumeId"`
	Time     string `json:"time"`
	School   string `json:"school"`
	Major    string `json:"major"`
	Degree   string `json:"degree"`
}
type ResumeProjectExperience struct {
	ID          int    `json:"id"`
	ResumeId    int    `json:"resumeId"`
	Time        string `json:"time"`
	ProjectInfo string `json:"projectInfo"`
}
type ResumeSkill struct {
	ID       int    `json:"id"`
	ResumeId int    `json:"resumeId"`
	Name     string `json:"name"`
	Percent  string `json:"percent"`
}
type ResumeWorkExperience struct {
	ID       int    `json:"id"`
	ResumeId int    `json:"resumeId"`
	Time     string `json:"time"`
	Position string `json:"position"`
	Company  string `json:"company"`
}
type Resume struct {
	ID                string                    `json:"id"`
	BaseInfo          ResumeBasicInfo           `json:"baseInfo"`
	CareerObjective   ResumeCareerObjective     `json:"careerObjective"`
	EducationInfo     []ResumeEducation         `json:"educationInfo"`
	ProjectExperience []ResumeProjectExperience `json:"projectExperience"`
	SkillInfo         []ResumeSkill             `json:"skills"`
	WorkExperience    []ResumeWorkExperience    `json:"workExperience"`
}
type ResumeIndex struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	UpdatedAt string `json:"update_time"`
}
