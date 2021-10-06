package service

import (
	"GinNaiveAdmin/global"
	"GinNaiveAdmin/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

// JsonInBlacklist 拉黑jwt
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error
func JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	err = global.GNA_DB.Create(&jwtList).Error
	return
}

// IsBlacklist 判断JWT是否在黑名单内部
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool
func IsBlacklist(jwt string) bool {
	isNotFound := errors.Is(global.GNA_DB.Where("jwt = ?", jwt).First(&model.JwtBlacklist{}).Error, gorm.ErrRecordNotFound)
	return !isNotFound
}

// GetRedisJWT 从redis取jwt
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: err error, redisJWT string
func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.GNA_REDIS.Get(userName).Result()
	return err, redisJWT
}

// SetRedisJWT jwt存入redis并设置过期时间
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: userName string
//@return: err error, redisJWT string
func SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Now().Add(time.Hour * time.Duration(global.GNA_CONF.JWT.AccessExpiresTime)).Unix()
	err = global.GNA_REDIS.Set(userName, jwt, time.Duration(timer)).Err()
	return err
}
