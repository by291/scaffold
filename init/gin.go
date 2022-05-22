package init

import (
	"by291.fun/scaffold/global"
	"github.com/gin-gonic/gin"
)

func GinInit() {
	conf := &global.Config

	// set gin mode
	var ginMode string
	if conf.Log.Mode == "dev" {
		ginMode = gin.DebugMode
	} else if conf.Log.Mode == "prod" {
		ginMode = gin.ReleaseMode
	}
	gin.SetMode(ginMode)
}
