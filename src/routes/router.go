package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jfxdev/petsaur/src/routes/health"
)

var router *gin.RouterGroup

func Register(engine *gin.Engine) {
	router = engine.Group("/")
	health.AddRoutes(router)
}

func Router() *gin.RouterGroup {
	return router
}
