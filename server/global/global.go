package global

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/config"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GvaDb        *gorm.DB
	GvaDbList    map[string]*gorm.DB
	GvaRedis     redis.UniversalClient
	GvaRedisList map[string]redis.UniversalClient
	GvaMongo     *qmgo.QmgoClient
	GvaConfig    config.Server
	GvaVp        *viper.Viper
	// GvaLog    *logging.Logger
	GvaLog                *zap.Logger
	GvaTimer              = timer.NewTimerTask()
	GvaConcurrencyControl = &singleflight.Group{}
	GvaRouters            gin.RoutesInfo
	GvaActiveDbname       *string
	BlackCache            local_cache.Cache
	lock                  sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GvaDbList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GvaDbList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

func GetRedis(name string) redis.UniversalClient {
	redis, ok := GvaRedisList[name]
	if !ok || redis == nil {
		panic(fmt.Sprintf("redis `%s` no init", name))
	}
	return redis
}
