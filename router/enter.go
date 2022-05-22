package router

import (
	"by291.fun/scaffold/global"
	"by291.fun/scaffold/middleware"
	"github.com/gin-gonic/gin"
)

// Routers init all and return a router
func Routers() *gin.Engine {
	conf := &global.Config
	r := gin.New()

	printStack := false
	if conf.Log.Mode == "dev" {
		printStack = true
	}

	// add all middlewares
	group := r.Group("/" + conf.App.Name)
	group.Use(
		middleware.Logger(),
		middleware.Recovery(printStack),
		middleware.Secure(),
		middleware.Cors(),
	)

	// init all routers
	// InitXXXRouter(group)

	return r
}
