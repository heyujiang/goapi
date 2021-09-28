package router

import (
	"github.com/gin-gonic/gin"
	"goapi/handler/sd"
	"goapi/handler/user"
	"goapi/router/middleware"
	"net/http"
)

func Load(g *gin.Engine,middlewares []gin.HandlerFunc) *gin.Engine{
	g.Use(gin.Recovery())

	g.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusOK,"The incorrect API route.")
	})

	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middlewares...)

	sdrg := g.Group("/sd")
	{
		sdrg.GET("/health",sd.HealthCheck)
		sdrg.GET("/disk",sd.DiskCheck)
		sdrg.GET("/cpu",sd.CPUCheck)
		sdrg.GET("/ram",sd.RAMCheck)
	}


	userrg := g.Group("/user")
	{
		userrg.POST("/create",user.Create)
	}


	return g
}