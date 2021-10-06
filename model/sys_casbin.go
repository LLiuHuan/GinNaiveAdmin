package model

type CasbinModel struct {
	Id    int    `json:"id" gorm:"column:id"`         // 主键
	PType string `json:"p_type" gorm:"column:p_type"` // Policy Type - 用于区分 policy和 group(role)
	V0    string `json:"v0" gorm:"column:v0"`         // subject   rolename
	V1    string `json:"v1" gorm:"column:v1"`         // object    path
	V2    string `json:"v2" gorm:"column:v2"`         // action    method
	V3    string // 这个和下面的字段无用，仅预留位置，如果你的不是
	V4    string // sub, obj, act的话才会用到
	V5    string // 如 sub, obj, act, suf就会用到 V3
}
