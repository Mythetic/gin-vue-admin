package internal

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
// Author [SliverHorn](https://github.com/SliverHorn)
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	var general config.GeneralDB
	switch global.GvaConfig.System.DbType {
	case "mysql":
		general = global.GvaConfig.Mysql.GeneralDB
	case "pgsql":
		general = global.GvaConfig.Pgsql.GeneralDB
	case "oracle":
		general = global.GvaConfig.Oracle.GeneralDB
	case "sqlite":
		general = global.GvaConfig.Sqlite.GeneralDB
	case "mssql":
		general = global.GvaConfig.Mssql.GeneralDB
	default:
		general = global.GvaConfig.Mysql.GeneralDB
	}
	return &gorm.Config{
		Logger: logger.New(NewWriter(general, log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      general.LogLevel(),
			Colorful:      true,
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
