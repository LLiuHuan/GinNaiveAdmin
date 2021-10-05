package utils

import (
	"GinNaiveAdmin/global"
	"os"
	"path"
	"time"

	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: GetWriteSyncer
//@description: zap logger 中加入 file-rotatelogs, 按照日期轮询创建日志（省的自己写了）
//@return: zapcore.WriteSyncer, error

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(global.GNA_CONF.Zap.Directory, "%Y-%m-%d.log"), // 创建路径
		zaprotatelogs.WithMaxAge(7*24*time.Hour),                 // 清除日志的最长时间
		zaprotatelogs.WithRotationTime(24*time.Hour),             // 切割日志的间隔，一天一个
	)
	if global.GNA_CONF.Zap.LogInConsole { // 同步在命令窗也显示
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
