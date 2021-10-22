package router

import (
	"github.com/gin-gonic/gin"
	"goapi/controller"
	"goapi/router/middleware"
	"net/http"
)

func Load(g *gin.Engine, middlewares []gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())

	g.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "The incorrect API route.")
	})

	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middlewares...)

	//服务器信息相关接口
	sdrg := g.Group("/server")
	{
		sdrg.GET("/health", controller.HealthCheck)
		sdrg.GET("/disk", controller.DiskCheck)
		sdrg.GET("/cpu", controller.CPUCheck)
		sdrg.GET("/ram", controller.RAMCheck)
	}

	userrg := g.Group("/user")
	userrg.Use(middleware.AuthMiddleware) //JWT 用户登录中间件
	{
		userrg.GET("/:id", controller.GetUserInfo)   //获取指定id的用户的详细信息
		userrg.POST("", controller.CreateUser)       //创建用户
		userrg.PUT("/:id", controller.UpdateUser)    //更新用户
		userrg.DELETE("/:id", controller.DeleteUser) //删除用户
		userrg.GET("", controller.UserList)          //用户列表
	}

	g.POST("/login", controller.Login) //登录

	g.POST("/createShortUrl", controller.GenerateShortUrl)
	g.GET("/:shortStr", controller.RedirectToLongUrl)

	course := g.Group("/course")
	{
		course.GET("", controller.CourseList)
		course.GET("/lessons/:id", controller.Lessons)
		course.GET("/lessonDetail/:id", controller.LessonDetail)
	}

	return g
}
