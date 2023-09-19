package resumeController

import (
	"net/http"
	"resume_backend/models"
	"resume_backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary 获取简历列表
// @Description 获取列表
// @Tags resume
// @Success 200{array} models.Resume
// @Router /api/resume/list  [get]
func GetResumeList(c *gin.Context) {
}

// @Summary 获取简历详情
// @Description 获取简历详情
// @Tags resume
// @Param id path int true "简历ID"
// @Success 200 {object} models.Resume
// @Router /api/resume/{id}  [get]
func GetSingleResume(c *gin.Context) {
}

// @Summary 创建简历
// @Description 创建简历
// @Tags resume
// @Param req body models.Resume true "简历信息"
// @Success 200 {string} string "ok"
// @Router /api/resume/add  [post]
func CreateResume(c *gin.Context) {
	// 从前端拿到参数
	var resumeInfo models.Resume
	if err := c.BindJSON(&resumeInfo); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		// 存储到数据库
		repositories.CreateResume(resumeInfo)
		c.IndentedJSON(http.StatusCreated, resumeInfo)
	}
}

// @Summary 编辑简历辑简历
// @Tags resume
// @Param req body models.Resume true "简历信息"
// @Router /api/resume/update  [post]
func UpdateResume(c *gin.Context) {
}

// @Summary 删除简历
// @Description 删除简历
// @Tags resume
// @Param id pah int  false "简历id""
// @Router /api/resume/{id}  [delete]
func DeleteResume(c *gin.Context) {
}
