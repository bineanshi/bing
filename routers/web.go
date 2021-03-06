package routers

import (
	. "SnatchBingWallpaper/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.LoadHTMLGlob("templates/*/*.tmpl")
	router.GET("/", List)
	router.GET("/refresh", Refresh)
	router.GET("/x/:date", List)


	return router
}
