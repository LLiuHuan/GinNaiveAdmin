package conf

type Server struct {
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	Casbin    Casbin    `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	Zap       Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql     Mysql     `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	RateLimit RateLimit `mapstructure:"rate-limit" json:"rate-limit" yaml:"rate-limit"`
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Captcha   Captcha   `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}
