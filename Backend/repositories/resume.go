package repositories

import (
	"resume_backend/models"

	_ "github.com/go-sql-driver/mysql"
)

func CreateResume(resumeInfo models.Resume) {
	// db := getDB()
}
func GetSingleResume(id string) (real string, err error) {
	data, err := id, nil
	return data, err
}
func GetResumeList() (list models.Resume, err error) {
	data, err := models.Resume{}, nil
	return data, err
}
func UpdateResume(resumeInfo models.Resume) {
	// db := getDB()
}
func DeleteResume(id string) error {
	// db := getDB()
	return nil
}
