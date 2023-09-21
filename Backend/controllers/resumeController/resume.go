package resumeController

import (
	"net/http"
	"resume_backend/consts"
	"resume_backend/models"
	"resume_backend/pkg/ctl"
	"resume_backend/pkg/utils"
	"resume_backend/service/resumeService"

	"github.com/gin-gonic/gin"
)

// ResumeCreateHandler @Tags RESUME
// @Summary 创建简历
// @Produce json
// @Accept json
// @Param     data    body    resumeService.resumeService    true      ""
// @Success 200 {object} serializer.ResponseUser "{"success":true,"data":{},"msg":"创建成功"}"
// @Failure 500 {object} serializer.ResponseUser "{"status":500,"data":{},"Msg":{},"Error":"error"}"
// @Router /resume/add  [post]
func ResumeCreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.CreateResumeReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := resumeService.GetResumeService()
			resp, err := l.CreateResume(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, ctl.ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			utils.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ctl.ErrorResponse(err))
		}
	}
}

// @Router /user/login [post]
// @Summary 获取简历列表
// @Description 获取列表
// @Tags resume
// @Success 200{array} models.Resume
// @Router /api/resume/list  [get]
func ResumeListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.ListResumeReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			if req.Limit == 0 {
				req.Limit = consts.BasePageSize
			}
			l := resumeService.GetResumeService()
			resp, err := l.ListResume(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ctl.ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			utils.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ctl.ErrorResponse(err))
		}
	}
}

// @Summary 获取简历详情
// @Description 获取简历详情
// @Tags resume
// @Param id path int true "简历ID"
// @Success 200 {object} models.Resume
// @Router /api/resume/{id}  [get]
func ResumeDetailHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.DetailReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := resumeService.GetResumeService()
			resp, err := l.DetailResume(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ctl.ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			utils.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ctl.ErrorResponse(err))
		}
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
