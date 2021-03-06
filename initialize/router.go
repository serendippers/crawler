package initialize

import (
	"crawler/global"
	"crawler/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "crawler/docs"
)

func Routers() *gin.Engine {
	engine := gin.Default()

	//Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.LOG.Debug("use middleware logger")

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Debug("use middleware swagger")

	apiGroup := engine.Group("")

	router.InitBaseRouter(apiGroup)

	global.LOG.Info("router register success")

	return engine
}
