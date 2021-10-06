package model

import (
	"GinNaiveAdmin/global"
)

type JwtBlacklist struct {
	global.GNA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
