package main

import (
	. "SnatchBingWallpaper/routers"
)
func main()  {
	router := InitRouter()
	router.Run(":8000")
}
