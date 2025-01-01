package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ragulmathawa/go-react-auth-app/pkg/utils"
)

func InitApi(router *gin.RouterGroup, appConfig utils.AppConfig) {
	router.GET("hello", utils.VerifySession(nil), func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
}
