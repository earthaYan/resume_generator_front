package router

import (
	"resume_backend/controllers/resumeController"

	"github.com/gin-gonic/gin"
)

func (r Router) Resume(g *gin.RouterGroup) {
	rg := g.Group("/resume")
	{
		// 查看简历列表
		rg.GET("/list", resumeController.GetResumeList)
		// 查询简历详情
		rg.GET("/:id", resumeController.GetSingleResume)
		// 添加简历
		rg.POST("/add", resumeController.CreateResume)
		// 修改简历
		rg.POST("/update", resumeController.UpdateResume)
		// 删除简历
		rg.DELETE("/:id", resumeController.DeleteResume)
	}
}
