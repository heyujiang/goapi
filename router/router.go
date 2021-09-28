package router

import (
	"github.com/gin-gonic/gin"
	"goapi/handler/sd"
	"goapi/handler/user"
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

	sdrg := g.Group("/sd")
	{
		sdrg.GET("/health", sd.HealthCheck)
		sdrg.GET("/disk", sd.DiskCheck)
		sdrg.GET("/cpu", sd.CPUCheck)
		sdrg.GET("/ram", sd.RAMCheck)
	}

	userrg := g.Group("/v1/user")
	{
		userrg.POST("", user.Create)       //创建用户
		userrg.DELETE("/:id", user.Delete) //删除用户
		userrg.PUT("/:id", user.Update)    //更新用户
		userrg.GET("", user.List)          //用户列表
		userrg.GET("/:id", user.Get)       //获取指定id的用户的详细信息
	}

	return g
}
