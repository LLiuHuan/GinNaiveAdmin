package global

import (
	"GinNaiveAdmin/conf"

	ut "github.com/go-playground/universal-translator"

	"github.com/go-redis/redis"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GNA_VP    *viper.Viper
	GNA_LOG   *zap.Logger
	GNA_DB    *gorm.DB
	GNA_CONF  conf.Server
	GNA_REDIS *redis.Client
	GNA_TRANS ut.Translator
)
