package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/router"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	public := engine.Group(global.GvaConfig.System.RouterPrefix).Group("")
	private := engine.Group(global.GvaConfig.System.RouterPrefix).Group("")
	private.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	router.Router.Info.Init(public, private)
}
