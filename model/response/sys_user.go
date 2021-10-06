package response

type LoginResponse struct {
	Username  string `json:"userName" gorm:"comment:用户登录名"`                                                 // 用户登录名 	// 用户登录密码
	NickName  string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                     // 用户昵称
	HeaderImg string `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"` // 用户头像

	AccessToken string `json:"access_token"`
	//RefreshToken string        `json:"refresh_token"`
	ExpiresAt int64 `json:"expiresAt"`
}
