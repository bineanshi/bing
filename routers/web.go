package routers

import (
	. "SnatchBingWallpaper/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	router.LoadHTMLGlob(filepath.Join(pwd,"templates/*/*.tmpl"))
	router.GET("/:date", List)

	return router
}
