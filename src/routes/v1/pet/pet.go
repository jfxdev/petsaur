package pet

import (
	"net/http"

	service "github.com/jfxdev/petsaur/src/services/v1/pet"

	"github.com/gin-gonic/gin"
)

func init() {
	svc = service.New()
}

var svc *service.Service

func setRootPath(router *gin.RouterGroup) *gin.RouterGroup {
	return router.Group("/pets")
}

func AddRoutes(router *gin.RouterGroup) {
	group := setRootPath(router)
	group.GET("/", pets)
	group.GET("/routines", routines)
}

func routines(c *gin.Context) {
	c.JSON(http.StatusOK, svc.Routines())
}

func pets(c *gin.Context) {
	response, err := svc.Get()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, response)
}
