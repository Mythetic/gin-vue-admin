package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func bizModel() error {
	db := global.GvaDb
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
