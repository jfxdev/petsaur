package v1

import (
	"fmt"

	"github.com/jfxdev/petsaur/src/routes/v1/crypto"
	"github.com/jfxdev/petsaur/src/routes/v1/pet"

	"github.com/gin-gonic/gin"
)

const version = "v1"

var router *gin.RouterGroup

func Register(engine *gin.Engine) {
	router = engine.Group(fmt.Sprintf("/%s", version))
	pet.AddRoutes(router)
	crypto.AddRoutes(router)
}

func Router() *gin.RouterGroup {
	return router
}
