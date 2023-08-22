package repositories

import "resume_backend/models"

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"resume_backend/models"
)

func CreateResume(resumeInfo models.Resume) {
	db := getDB()
}
