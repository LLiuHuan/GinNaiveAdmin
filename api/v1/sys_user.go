package v1

import (
	"GinNaiveAdmin/global"
	"GinNaiveAdmin/middlewares"
	"GinNaiveAdmin/model"
	"GinNaiveAdmin/model/request"
	"GinNaiveAdmin/model/response"
	"GinNaiveAdmin/service"
	"GinNaiveAdmin/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Login 登录方法
// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var L request.Login

	if errStr, err := utils.BaseValidator(&L, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	if L.CaptchaId != "" && L.Captcha != "" {
		if !VerifyCaptcha(L.CaptchaId, L.Captcha) {
			response.FailWithMessage("验证码错误", c)
			return
		}
	}

	U := &model.SysUser{Username: L.Username, Password: L.Password}
	if err, user := service.Login(U); err != nil {
		global.GNA_LOG.Error("登陆失败! 用户名不存在或者密码错误", zap.Any("err", err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
		return
	} else {
		jwtUser := model.JwtUser{
			UUID:        user.UUID,
			Username:    user.Username,
			NickName:    user.NickName,
			HeaderImg:   user.HeaderImg,
			AuthorityId: user.AuthorityId,
		}

		tokenNext(c, jwtUser)
	}
}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user model.JwtUser) {
	j := middlewares.NewJWT() // 唯一签名

	claims := request.CustomClaims{
		UserInfo: user,
		//BufferTime: global.GEA_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),                                                                       // 签名生效时间
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(global.GNA_CONF.JWT.AccessExpiresTime)).Unix(), // 过期时间 配置文件
			Issuer:    "LLiuHuan",                                                                              // 签名的发行者
		},
	}

	accessToken, err := j.CreateToken(claims)
	if err != nil {
		global.GNA_LOG.Error("获取token失败", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	// 多点登录
	if !global.GNA_CONF.System.UseMultipoint {
		response.OkWithDetailed(response.LoginResponse{
			Username:    user.Username,
			NickName:    user.NickName,
			HeaderImg:   user.HeaderImg,
			AccessToken: accessToken,
			//RefreshToken: refreshToken,
			ExpiresAt: claims.StandardClaims.ExpiresAt,
		}, "登录成功", c)
		return
	}
	// 单点登录需要使用redis处理token
	if err, jwtStr := service.GetRedisJWT(user.Username); err == redis.Nil {
		if err := service.SetRedisJWT(accessToken, user.Username); err != nil {
			global.GNA_LOG.Error("设置登录状态失败", zap.Any("err", err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			Username:    user.Username,
			NickName:    user.NickName,
			HeaderImg:   user.HeaderImg,
			AccessToken: accessToken,
			//RefreshToken: refreshToken,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GNA_LOG.Error("设置登录状态失败", zap.Any("err", err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := service.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := service.SetRedisJWT(accessToken, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			Username:    user.Username,
			NickName:    user.NickName,
			HeaderImg:   user.HeaderImg,
			AccessToken: accessToken,
			//RefreshToken: refreshToken,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// 从Gin的Context中获取从jwt解析出来的用户ID
func getUserID(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		global.GNA_LOG.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件")
		return uuid.Nil
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.UserInfo.UUID
	}
}
