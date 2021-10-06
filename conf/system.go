package conf

// System 基础配置
type System struct {
	Name          string `mapstructure:"name" json:"name" yaml:"name"`                              // 项目名称
	Mode          string `mapstructure:"mode" json:"mode" yaml:"mode"`                              // 项目模式
	Port          int    `mapstructure:"port" json:"port" yaml:"port"`                              // 项目使用端口
	Version       string `mapstructure:"version" json:"version" yaml:"version"`                     // 项目版本
	StartTime     string `mapstructure:"start_time" json:"start_time" yaml:"start_time"`            // 项目开始时间
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // 多点登录拦截
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                 // 环境值   目前没什么用
}
