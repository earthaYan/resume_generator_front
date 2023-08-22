package main

import (
	"reflect"
	docs "resume_backend/docs"
	"resume_backend/repositories"
	"resume_backend/router"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// 初始化数据库
	repositories.InitMysql()
	// 初始化gin
	r := gin.Default()
	router := r.Group("/api")
	ReadRouters(router)
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":3000")
}

// ReadRouters 读取router下的路由组
func ReadRouters(g *gin.RouterGroup) {
	routes := router.Router{}
	val := reflect.ValueOf(routes)
	// 获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	for i := 0; i < numOfMethod; i++ {
		// 断言特定类型的方法
		fn, ok := val.Method(i).Interface().(func(g *gin.RouterGroup))
		if !ok {
			continue
		}
		fn(g)
	}
}
