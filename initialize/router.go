package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var r = gin.Default()

	PublicGroupV1 := r.Group("/v1")
	{
		PublicGroupV1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})
	}

	return r
}
