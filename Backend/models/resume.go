package models

type ResumeBasicInfo struct {
	ID              int    `json:"base_info_id"`
	ResumeId        int    `json:"resume_id"`
	Name            string `json:"name"`
	Age             int    `json:"age"`
	ExperienceYears int    `json:"experience_years"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	BriefInfo       string `json:"brief_info"`
}
type ResumeCareerObjective struct {
	ID       int    `json:"career_objective_id"`
	ResumeId int    `json:"resume_id"`
	City     string `json:"city"`
	Salary   string `json:"salary"`
	Status   string `json:"status"`
}
type ResumeEducation struct {
	ID       int    `json:"education_id"`
	ResumeId int    `json:"resume_id"`
	Time     string `json:"time"`
	School   string `json:"school"`
	Major    string `json:"major"`
	Degree   string `json:"degree"`
}
type ResumeProjectExperience struct {
	ID          int    `json:"project_exerience_id"`
	ResumeId    int    `json:"resume_id"`
	ProjectInfo string `json:"project_info"`
}
type ResumeSkill struct {
	ID       int    `json:"skill_id"`
	ResumeId int    `json:"resume_id"`
	Name     string `json:"name"`
	Percent  string `json:"percent"`
}
type ResumeWorkExperience struct {
	ID       int    `json:"work_experience_id"`
	ResumeId int    `json:"resume_id"`
	Time     string `json:"time"`
	Position string `json:"position"`
	Company  string `json:"company"`
}
type Resume struct {
	ID                int                       `json:"id"`
	Name              string                    `json:"name"`
	BaseInfo          ResumeBasicInfo           `json:"base_info"`
	CareerObjective   []ResumeCareerObjective   `json:"career_objective"`
	EducationInfo     []ResumeEducation         `json:"education_info"`
	ProjectExperience []ResumeProjectExperience `json:"project_experience"`
	SkillInfo         []ResumeSkill             `json:"skil_info"`
	WorkExperience    []ResumeWorkExperience    `json:"work_experience"`
}
