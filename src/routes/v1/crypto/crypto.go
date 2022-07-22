package crypto

import (
	"net/http"

	service "github.com/jfxdev/petsaur/src/services/v1/crypto"

	"github.com/gin-gonic/gin"
)

func init() {
	svc = service.New()
}

var svc *service.Service

func setRootPath(router *gin.RouterGroup) *gin.RouterGroup {
	return router.Group("/crypto")
}

func AddRoutes(router *gin.RouterGroup) {
	group := setRootPath(router)
	group.GET("/", list)
}

func list(c *gin.Context) {
	c.JSON(http.StatusOK, svc.Crypto())
}
