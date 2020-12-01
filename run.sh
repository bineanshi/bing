#!/bin/bash
killall -9 SnatchBingWallpaper
killall -9 BingWallpaperServer
nohup ./SnatchBingWallpaper &
nohup ./BingWallpaperServer &