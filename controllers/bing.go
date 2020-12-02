package controllers

import (
	. "SnatchBingWallpaper/database"
	. "SnatchBingWallpaper/lib"
	. "SnatchBingWallpaper/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func List(c *gin.Context) {
	lastDay, _ := time.ParseInLocation("20060102", "20201125", time.Local)
	timeStr := time.Now().AddDate(0,0,-2).Format("20060102")
	currentDay, _ := time.Parse("20060102", timeStr)
	date, _ := time.ParseInLocation("20060102", c.Param("date"), time.Local)
	StartDate := string(date.Format("20060102"))
	var wallpaper Wallpaper
	DB.Find(&wallpaper, "start_date = ?" , StartDate)

	if wallpaper.ID < 1 {
		if date.Before(lastDay) {
			DB.Order("start_date asc").First(&wallpaper)
		}else if date.After(currentDay){
			DB.Order("start_date asc").Last(&wallpaper)
		}

	}
	pageDate,_ := time.ParseInLocation("20060102", wallpaper.StartDate, time.Local)
	nextDay := pageDate.AddDate(0,0,1)
	previousDay := pageDate.AddDate(0,0,-1)

	c.HTML(http.StatusOK, "bing/index.tmpl", gin.H{
		"image": wallpaper,
		"previous_start_date": string(previousDay.Format("20060102")),
		"next_start_date": string(nextDay.Format("20060102")),
	})
}

func Refresh(c *gin.Context) {
	CreateDataDir("./", "download_tmp")

	Snatch()

	//RemoveDataDir("./","download_tmp")

	var wallpaper Wallpaper
	DB.Order("start_date asc").Last(&wallpaper)
	pageDate,_ := time.ParseInLocation("20060102", wallpaper.StartDate, time.Local)
	nextDay := pageDate.AddDate(0,0,1)
	previousDay := pageDate.AddDate(0,0,-1)
	c.HTML(http.StatusOK, "bing/index.tmpl", gin.H{
		"image": wallpaper,
		"previous_start_date": string(previousDay.Format("20060102")),
		"next_start_date": string(nextDay.Format("20060102")),
	})
}