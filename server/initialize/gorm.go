package initialize

import (
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GvaConfig.System.DbType {
	case "mysql":
		global.GvaActiveDbname = &global.GvaConfig.Mysql.Dbname
		return GormMysql()
	case "pgsql":
		global.GvaActiveDbname = &global.GvaConfig.Pgsql.Dbname
		return GormPgSql()
	case "oracle":
		global.GvaActiveDbname = &global.GvaConfig.Oracle.Dbname
		return GormOracle()
	case "mssql":
		global.GvaActiveDbname = &global.GvaConfig.Mssql.Dbname
		return GormMssql()
	case "sqlite":
		global.GvaActiveDbname = &global.GvaConfig.Sqlite.Dbname
		return GormSqlite()
	default:
		global.GvaActiveDbname = &global.GvaConfig.Mysql.Dbname
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.GvaDb
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysIgnoreApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCodePackage{},
		system.SysExportTemplate{},
		system.Condition{},
		system.JoinTemplate{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},
	)
	if err != nil {
		global.GvaLog.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	err = bizModel()

	if err != nil {
		global.GvaLog.Error("register biz_table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GvaLog.Info("register table success")
}
