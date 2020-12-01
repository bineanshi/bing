package controllers

import (
	. "SnatchBingWallpaper/database"
	. "SnatchBingWallpaper/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func List(c *gin.Context) {
	var StartDate string
	date := c.Param("date")
	dateNum, _ := strconv.Atoi(date)
	timestamp := time.Now().AddDate(0, 0, -1)
	today := string(timestamp.Format("20060102"))
	todayNum,_ := strconv.Atoi(today)
	if dateNum < 20201124 || dateNum >= todayNum{
		dateNum = todayNum
	}
	StartDate = strconv.Itoa(dateNum)
	var wallpaper Wallpaper
	DB.Find(&wallpaper, "start_date = ?" , StartDate)
	c.HTML(http.StatusOK, "bing/index.tmpl", gin.H{
		"image": wallpaper,
		"previous_start_date": dateNum - 1,
		"next_start_date": dateNum + 1,
	})
}