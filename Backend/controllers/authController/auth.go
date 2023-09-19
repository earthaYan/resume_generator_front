package authcontroller

import (
	"net/http"
	"resume_backend/models"
	"resume_backend/pkg/ctl"
	"resume_backend/pkg/utils"
	authservice "resume_backend/service/authService"

	"github.com/gin-gonic/gin"
)

// UserRegisterHandler @Tags USER
// @Summary 用户注册
// @Produce json
// @Accept json
// @Param data body service.UserService true "用户名, 密码"
// @Success 200 {object} serializer.ResponseUser "{"status":200,"data":{},"msg":"ok"}"
// @Failure 500  {object} serializer.ResponseUser "{"status":500,"data":{},"Msg":{},"Error":"error"}"
// @Router /user/register [post]
func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.UserServiceRegisterReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := authservice.GetUserService()
			resp, err := l.UserRegister(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ctl.ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			utils.LogrusObj.Info(err)
			ctx.JSON(http.StatusBadRequest, ctl.ErrorResponse(err))
		}
	}
}

// UserLoginHandler @Tags USER
// @Summary 用户登录
// @Produce json
// @Accept json
// @Param     data    body     service.UserService    true      "user_name, password"
// @Success 200 {object} serializer.ResponseUser "{"success":true,"data":{},"msg":"登陆成功"}"
// @Failure 500 {object} serializer.ResponseUser "{"status":500,"data":{},"Msg":{},"Error":"error"}"
// @Router /user/login [post]
func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.UserServiceLoginReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := authservice.GetUserService()
			resp, err := l.UserLogin(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ctl.ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			utils.LogrusObj.Info(err)
			ctx.JSON(http.StatusBadRequest, ctl.ErrorResponse(err))
		}
	}
}
