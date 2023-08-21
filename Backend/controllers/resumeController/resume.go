package resumeController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"resume_backend/models"
)

// @Summary 获取简历列表
// @Description 获取简历列表
// @Tags resume
// @Success 200 {array} models.Resume
// @Router /api/resume/list  [get]
func GetResumeList(c *gin.Context) {
	fmt.Println(&models.Resume{})
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
}

// @Summary 编辑简历
// @Description 编辑简历
// @Tags resume
// @Param req body models.Resume true "简历信息"
// @Success 200 {string} string "ok"
// @Router /api/resume/update  [post]
func UpdateResume(c *gin.Context) {

}

// @Summary 删除简历
// @Description 删除简历
// @Tags resume
// @Param id path int  false "简历id"
// @Success 200 {string} string "ok"
// @Router /api/resume/{id}  [delete]
func DeleteResume(c *gin.Context) {

}
