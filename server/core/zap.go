package core

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/core/internal"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GvaConfig.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.GvaConfig.Zap.Director)
		_ = os.Mkdir(global.GvaConfig.Zap.Director, os.ModePerm)
	}
	levels := global.GvaConfig.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger = zap.New(zapcore.NewTee(cores...))
	if global.GvaConfig.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
