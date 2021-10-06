package initialize

import (
	"GinNaiveAdmin/global"
	"GinNaiveAdmin/middlewares"
	"GinNaiveAdmin/router"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	var r = gin.Default()
	global.GNA_LOG.Info("use middleware logger")

	// 跨域
	r.Use(middlewares.Cors())
	// 令牌桶 限流
	r.Use(middlewares.RateLimitMiddleware(time.Second))
	// IP 限流
	if global.GNA_CONF.RateLimit.IpVerify {
		r.Use(middlewares.IpVerifyMiddleware())
	}

	// swagger
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	global.GNA_LOG.Info("register swagger handler")

	r.Use(middlewares.OperationRecord())
	PublicGroupV1 := r.Group("/v1")
	{
		router.InitBaseRouter(PublicGroupV1)
	}

	PrivateGroupV1 := r.Group("/v1")
	PrivateGroupV1.Use(middlewares.JWTAuto()).Use(middlewares.CasbinHandler())
	{
		PrivateGroupV1.GET("pong", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "测试1",
				"data": "ping",
			})
		})
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
