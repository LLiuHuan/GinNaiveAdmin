package middlewares

import (
	"GinNaiveAdmin/global"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/juju/ratelimit"
)

// RateLimitMiddleware 令牌桶限流
func RateLimitMiddleware(fillInterval time.Duration) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, global.GNA_CONF.RateLimit.Cap, global.GNA_CONF.RateLimit.Quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.String(http.StatusForbidden, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}

// IpVerifyMiddleware IP限流
func IpVerifyMiddleware() gin.HandlerFunc {
	blacklist := make(map[string]int64, 0)
	return func(c *gin.Context) {
		//visitorIP := ctx.Request.Header.Get("X-real-ip")
		visitorIP := c.ClientIP()
		//fmt.Println("IP:        ", visitorIP)
		// 判断是否在黑名单
		timeOrigin := global.GNA_REDIS.HGet(global.GNA_CONF.RateLimit.IpListKey, visitorIP).Val()
		err := blackListVerify(timeOrigin, visitorIP, global.GNA_REDIS, blacklist[visitorIP])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		// 如果不存在于黑名单
		lenList, _ := global.GNA_REDIS.LLen(visitorIP).Result()
		// 如果第一次登陆或频率限制外
		if lenList == 0 {
			// 如果为空跳过，设置后跳过
			// 添加IP list
			global.GNA_REDIS.LPush(visitorIP, visitorIP)
			// 设置过期时间
			global.GNA_REDIS.Expire(visitorIP, time.Second)
			c.Next()
		} else if lenList > 0 && lenList < global.GNA_CONF.RateLimit.IpLimitCon {
			global.GNA_REDIS.LPush(visitorIP, visitorIP)
			c.Next()
		} else {
			if blacklist[visitorIP] == 0 {
				blacklist[visitorIP] = 10
			} else {
				blacklist[visitorIP] *= 2
			}
			// 加入黑名单
			global.GNA_REDIS.HSet(global.GNA_CONF.RateLimit.IpListKey, visitorIP, time.Now().Local().Unix())
			c.Abort()
			return
		}
	}
}

func blackListVerify(ot, visitorIP string, rdb *redis.Client, limitTime int64) error {
	// 如果有值，进一步判断
	if ot != "" {
		// 如果value的时间和当前时间差超过十分钟，解除限制
		timeOriginInt, _ := strconv.Atoi(ot)
		oTimeUnix := time.Unix(int64(timeOriginInt), 0).Local()
		subTime := time.Now().Sub(oTimeUnix)
		if subTime > time.Duration(limitTime)*time.Minute {
			// 超过限制时间 解除限制
			rdb.HDel(global.GNA_CONF.RateLimit.IpListKey, visitorIP)
			return nil
		} else {
			//return errors.New(fmt.Sprintf("您已被加入黑名单，剩余限制时间:%v", 10*time.Minute-subTime))
			return errors.New(fmt.Sprintf("请求速度过快，该IP已被限制，剩余限制时间:%v", time.Duration(limitTime)*time.Minute-subTime))
		}
	}
	// 不存在返回nil
	return nil
}
