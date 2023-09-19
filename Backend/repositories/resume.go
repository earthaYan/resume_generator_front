package repositories

import (
	"resume_backend/models"

	_ "github.com/go-sql-driver/mysql"
)

func CreateResume(resumeInfo models.Resume) error {
	return nil
}
func GetSingleResume(id string) error {
	return nil
}

func GetResumeList() error {
	return nil
}
func UpdateResume(resumeInfo models.Resume) error {
	return nil
}
func DeleteResume(id string) error {
	return nil
}
