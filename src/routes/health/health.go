package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRootPath(router *gin.RouterGroup) *gin.RouterGroup {
	return router.Group("/health")
}

func AddRoutes(router *gin.RouterGroup) {
	group := setRootPath(router)
	group.GET("/", health)
}

func health(c *gin.Context) {
	c.Status(http.StatusOK) 
}