package ast

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"path/filepath"
)

func init() {
	global.GvaConfig.AutoCode.Root, _ = filepath.Abs("../../../")
	global.GvaConfig.AutoCode.Server = "server"
}
