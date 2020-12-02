package models

import (
	"bytes"
	"encoding/json"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	. "SnatchBingWallpaper/database"
	. "SnatchBingWallpaper/lib"
)

func init()  {
	// Migrate the schema
	DB.AutoMigrate(&Wallpaper{})
}
type Wallpapers struct {
	Images []Wallpaper `json:images`
}
type Wallpaper struct{
	gorm.Model
	ID        int    `gorm:"primary_key"`
	StartDate string  `json:startdate,gorm:"size:255;index:start_date,unique"`
	FullStartDate string `json:fullstartdate`
	EndDate string `json:enddate`
	Url string   `json:url`
	UrlBase string `json:urlbase`
	Copyright string `json:copyright`
	CopyrightLink string `json:copyrightlink`
	Quiz string `json:quiz`
	Wp string `json:wp`
	QiNiu bool
}


func (wallpaper Wallpaper)DownloadImage()(bool,int) {
	var image Wallpaper
	DB.Where("full_start_date = ?",  wallpaper.FullStartDate).First(&image)
	if image.FullStartDate == "" {
		imagPath := "https://cn.bing.com" + wallpaper.Url
		resp, _ := http.Get(imagPath)
		body, _ := ioutil.ReadAll(resp.Body)
		regName := regexp.MustCompile(`[^\(| |,|，]*`)
		name := regName.FindString(wallpaper.Copyright)
		out, _ := os.Create("./download_tmp/" + name + ".jpg")
		io.Copy(out, bytes.NewReader(body))
		result := DB.Create(&wallpaper)
		log.Println(result)
		return true,wallpaper.ID
	} else {
		return false, 0
	}

}

func (wallpaper Wallpaper)UploadQiNiu()(bool)  {
	regName := regexp.MustCompile(`[^\(| |,|，]*`)
	name := regName.FindString(wallpaper.Copyright)
	imagPath := "./download_tmp/" + name + ".jpg"
	fileName := wallpaper.FullStartDate
	ok, _ := QiNiuUpload(imagPath,fileName)
	return ok
}

func Snatch() {

	body := GetBingUrl(0,7)

	var images Wallpapers
	json.Unmarshal([]byte(body), &images)
	for _,wallpaper := range images.Images{
		isCreate,id := wallpaper.DownloadImage()
		if isCreate {
			ok := wallpaper.UploadQiNiu()
			if ok {
				DB.Model(&Wallpaper{}).Where("id = ?", id).Update("QiNiu", true)
			} else {
				DB.Model(&Wallpaper{}).Where("id = ?", id).Update("QiNiu", false)
			}
		}
	}
}

func (wallpaper Wallpaper) GetUrl() string {
	var url string
	if wallpaper.QiNiu{
		url = "http://qiniu.g-bill.club/" + wallpaper.FullStartDate + ".jpg"
	}else{
		url = "https://cn.bing.com" + wallpaper.Url
	}
	return url
}