package router

import (
	"crawler/api"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(apiGroup *gin.RouterGroup) {
	baseRouter := apiGroup.Group("base")

	{
		baseRouter.GET("show", api.Show)
	}
}
