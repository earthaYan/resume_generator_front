package router

import (
	authcontroller "resume_backend/controllers/authController"
	resumeController "resume_backend/controllers/resumeController"
	"resume_backend/docs"
	"resume_backend/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		v1.POST("user/register", authcontroller.UserRegisterHandler())
		v1.POST("user/login", authcontroller.UserLoginHandler())
		authd := v1.Group("/")
		authd.Use(middleware.JWT())
		{
			authd.POST("resume/create", resumeController.ResumeCreateHandler())
		}
	}
	return r
}
