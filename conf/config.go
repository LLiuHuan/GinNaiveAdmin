package conf

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Casbin Casbin `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
}
