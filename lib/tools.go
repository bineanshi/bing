package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

const URL = "https://cn.bing.com/HPImageArchive.aspx"



func CreateDataDir(basePath string,folderName string) {
	folderPath := filepath.Join(basePath, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		os.Mkdir(folderPath, 0777)
		// 再修改权限
		os.Chmod(folderPath, 0777)
		fmt.Println(folderPath,"已创建")
	}
}

func RemoveDataDir(basePath string,folderName string) {
	folderPath := filepath.Join(basePath, folderName)
	err := os.RemoveAll(folderPath)
	if  err != nil {
		fmt.Println(folderPath,"未删除")
	}else{
		fmt.Println(folderPath,"已删除")
	}
}

func GetBingUrl(idx int,n int) (result string){
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("format", "hp")
	q.Add("idx", strconv.Itoa(idx))
	q.Add("n", strconv.Itoa(n))
	req.URL.RawQuery = q.Encode()

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	reg := regexp.MustCompile(`{.*}`)
	resBody := string(body)
	result = reg.FindString(resBody)
	return result
}
