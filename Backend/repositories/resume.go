package repositories

import (
	"log"
	"resume_backend/models"

	_ "github.com/go-sql-driver/mysql"
)

func CreateResume(resumeInfo models.Resume) {
	db := getDB()
	// 先向resume表插入简历名称-来源于基础信息的名称
	// 依次插入对应的关联表
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	sqlStr1 := `insert into resume(name) values(?)`
	result, execErr := tx.Exec(sqlStr1, resumeInfo.BaseInfo.Name)
	resume_id, _ := result.LastInsertId()
	base := resumeInfo.BaseInfo
	sqlStr2 := `INSERT INTO resume_basic_info(resume_id,name, age, experience_years, email, phone, brief_intro) VALUES (?,?, ?, ?,?, ?, ?)`
	_, execErr = tx.Exec(sqlStr2, resume_id, base.Name, base.Age, base.ExperienceYears, base.Email, base.Phone, base.BriefInfo)
	target := resumeInfo.CareerObjective
	sqlStr3 := `INSERT INTO resume_career_objective(resume_id, position, city, salary, status) VALUES (?, ?, ?, ?, ?)`
	_, execErr = tx.Exec(sqlStr3, resume_id, target.Position, target.City, target.Salary, target.Status)
	educations := resumeInfo.EducationInfo
	for _, education := range educations {
		sqlStr4 := `INSERT INTO resume_education_experience(resume_id, time, school,  degree,major) VALUES (?, ?, ?, ?,"")`
		_, execErr = tx.Exec(sqlStr4, resume_id, education.Time, education.School, education.Degree)
	}
	projects := resumeInfo.ProjectExperience
	for _, project := range projects {
		sqlStr5 := `INSERT INTO resume_project_experience(resume_id, time, project_info) VALUES (?, ?, ?)`
		_, execErr = tx.Exec(sqlStr5, resume_id, project.Time, project.ProjectInfo)
	}
	skills := resumeInfo.SkillInfo
	for _, skill := range skills {
		sqlStr6 := `INSERT INTO resume_skills(resume_id, name, percent) VALUES (?, ?, ?)`
		_, execErr = tx.Exec(sqlStr6, resume_id, skill.Name, skill.Percent)
	}
	works := resumeInfo.WorkExperience
	for _, work := range works {
		sqlStr7 := `INSERT INTO resume_work_experience(resume_id, time,company,position) VALUES (?, ?, ?,?)`
		_, execErr = tx.Exec(sqlStr7, resume_id, work.Time, work.Company, work.Position)
	}
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

}
func GetSingleResume(id string) (models.Resume, error) {
	db := getDB()
	var data = models.Resume{}
	sqlStr := `select * from resume_basic_info where resume_id=?`
	row := db.QueryRow(sqlStr, id)
	var (
		ID              int
		ResumeId        int
		Name            string
		Age             int
		ExperienceYears int
		Email           string
		Phone           string
		BriefInfo       string
	)
	err := row.Scan(&ID, &ResumeId, &Name, &Age, &ExperienceYears, &Email, &Phone, &BriefInfo)
	data.BaseInfo = models.ResumeBasicInfo{
		ID:              ID,
		ResumeId:        ResumeId,
		Name:            Name,
		Age:             Age,
		ExperienceYears: ExperienceYears,
		Email:           Email,
		Phone:           Phone,
		BriefInfo:       BriefInfo,
	}
	sqlStr = `select * from resume_career_objective where resume_id=?`
	var (
		Position string
		City     string
		Salary   string
		Status   string
	)
	row = db.QueryRow(sqlStr, id)
	err = row.Scan(&ID, &ResumeId, &Position, &City, &Salary, &Status)
	data.CareerObjective = models.ResumeCareerObjective{
		ID:       ID,
		ResumeId: ResumeId,
		Position: Position,
		City:     City,
		Salary:   Salary,
		Status:   Status,
	}

	sqlStr = `select * from resume_education_experience where resume_id=?`
	var educations []models.ResumeEducation
	rows, err := db.Query(sqlStr, id)
	for rows.Next() {
		var (
			ID       int
			ResumeId int
			Time     string
			School   string
			Degree   string
			Major    string
		)
		if err := rows.Scan(&ID, &ResumeId, &Time, &School, &Major, &Degree); err != nil {
			log.Fatal(err)
		}
		educations = append(educations, models.ResumeEducation{ID: ID, ResumeId: ResumeId, School: School, Time: Time, Degree: Degree, Major: Major})
	}
	data.EducationInfo = educations

	sqlStr = `select * from resume_project_experience where resume_id=?`
	var projects []models.ResumeProjectExperience
	rows, err = db.Query(sqlStr, id)
	for rows.Next() {
		var (
			ID          int
			ResumeId    int
			Time        string
			ProjectInfo string
		)
		if err := rows.Scan(&ID, &ResumeId, &ProjectInfo, &Time); err != nil {
			log.Fatal(err)
		}
		projects = append(projects, models.ResumeProjectExperience{ID: ID, ResumeId: ResumeId, Time: Time, ProjectInfo: ProjectInfo})
	}
	data.ProjectExperience = projects

	sqlStr = `select * from resume_skills where resume_id=?`
	var skills []models.ResumeSkill
	rows, err = db.Query(sqlStr, id)
	for rows.Next() {
		var (
			ID       int
			ResumeId int
			Name     string
			Percent  string
		)
		if err := rows.Scan(&ID, &ResumeId, &Name, &Percent); err != nil {
			log.Fatal(err)
		}
		skills = append(skills, models.ResumeSkill{ID: ID, ResumeId: ResumeId, Name: Name, Percent: Percent})
	}
	data.SkillInfo = skills

	sqlStr = `select * from resume_work_experience where resume_id=?`
	var works []models.ResumeWorkExperience
	rows, err = db.Query(sqlStr, id)
	for rows.Next() {
		var (
			ID       int
			ResumeId int
			Time     string
			Company  string
			Position string
		)
		if err := rows.Scan(&ID, &ResumeId, &Time, &Company, &Position); err != nil {
			log.Fatal(err)
		}
		works = append(works, models.ResumeWorkExperience{ID: ID, ResumeId: ResumeId, Time: Time, Position: Position, Company: Company})
	}
	data.WorkExperience = works
	data.ID = id
	return data, err

}

func GetResumeList() ([]models.ResumeIndex, error) {
	db := getDB()
	rows, err := db.Query(`select id, name,updated_at from resume`)
	if err != nil {
		log.Fatal(err)
	}
	var data []models.ResumeIndex
	for rows.Next() {
		var name string
		var id int
		var updated_at string
		if err := rows.Scan(&id, &name, &updated_at); err != nil {
			log.Fatal(err)
		}
		current := models.ResumeIndex{ID: id, Name: name, UpdatedAt: updated_at}
		data = append(data, current)
	}
	return data, err
}
func UpdateResume(resumeInfo models.Resume) {
	// resume_id := resumeInfo.ID

}
func DeleteResume(id string) error {
	db := getDB()
	tx, err := db.Begin()
	sqlStr1 := `DELETE  from resume where id=?`
	_, err = db.Exec(sqlStr1, id)
	sqlStr2 := `DELETE  from resume_basic_info where resume_id=?`
	_, err = db.Exec(sqlStr2, id)
	sqlStr3 := `DELETE  from resume_career_objective where resume_id=?`
	_, err = db.Exec(sqlStr3, id)
	sqlStr4 := `DELETE  from resume_education_experience where resume_id=?`
	_, err = db.Exec(sqlStr4, id)
	sqlStr5 := `DELETE  from resume_project_experience where resume_id=?`
	_, err = db.Exec(sqlStr5, id)
	sqlStr6 := `DELETE  from resume_skills where resume_id=?`
	_, err = db.Exec(sqlStr6, id)
	sqlStr7 := `DELETE from  resume_work_experience where resume_id=?`
	_, err = db.Exec(sqlStr7, id)
	if err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	return err
}
