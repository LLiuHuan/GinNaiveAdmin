package global

import (
	"GinNaiveAdmin/conf"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GNA_VP   *viper.Viper
	GNA_LOG  *zap.Logger
	GNA_DB   *gorm.DB
	GNA_CONF conf.Server
)
