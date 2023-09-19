package main

import (
	"resume_backend/pkg/utils"
	"resume_backend/repositories"
	"resume_backend/router"
)

func main() {
	loading()
	// 初始化gin
	r := router.NewRouter()
	r.Run(":3000")
}
func loading() {
	// 初始化数据库
	repositories.InitMysql()
	utils.InitLog()
}
