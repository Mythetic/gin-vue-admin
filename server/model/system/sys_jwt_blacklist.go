package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type JwtBlacklist struct {
	global.GvaModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
