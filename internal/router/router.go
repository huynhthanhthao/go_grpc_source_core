package router

import (
	"hrm/ent"

	"github.com/gin-gonic/gin"
)

func SetupRouter(client *ent.Client) *gin.Engine {
	r := gin.Default()

	return r
}
