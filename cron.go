package main

import (
	. "SnatchBingWallpaper/lib"
	. "SnatchBingWallpaper/models"
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func main()  {
	c := cron.New()
	c.AddFunc("0 1 16 * * *", func() {
		CreateDataDir("./", "download_tmp")

		Snatch()

		RemoveDataDir("./","download_tmp")
	})
	c.Start()

	t1 := time.NewTimer(time.Second * 10) // ?time.Second * 10 啥意思？ *100行吗？
	for {
		select {
		case <-t1.C:
			fmt.Println("Time now:", time.Now().Format("2006-01-02 15:04:05")) // 为何要专门制定这个时间
			t1.Reset(time.Second * 10)
		}
	}

}
