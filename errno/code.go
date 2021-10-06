package errno

import "net/http"

var (
	OK                  = &Errno{Code: http.StatusOK, Message: "OK"}
	InternalServerError = &Errno{Code: http.StatusInternalServerError, Message: "InternalServerError"}

	ParameterError      = &Errno{Code: 10100, Message: "参数错误"}
	ErrLoginParameter   = &Errno{Code: 10101, Message: "用户名或密码无法解析。"}
	ErrPasswordGenerate = &Errno{Code: 10102, Message: "生成密码出错"}
	ErrAbsent           = &Errno{Code: 10103, Message: "token absent"}                              // 令牌不存在
	ErrInvalid          = &Errno{Code: 10104, Message: "token invalid"}                             // 令牌无效
	ErrExpired          = &Errno{Code: 10105, Message: "token expired"}                             // 令牌过期
	ErrOther            = &Errno{Code: 10106, Message: "other error"}                               // 其他错误
	ErrAuthorization    = &Errno{Code: 10107, Message: "header Authorization has not Bearer token"} // header 有问题
	ErrNoPermission     = &Errno{Code: 10108, Message: "无权限"}                                       // 无权限

	ErrLoginNotExist = &Errno{Code: 20101, Message: "用户不存在"}
	ErrPasswordError = &Errno{Code: 20102, Message: "密码错误"}
)
