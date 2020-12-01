#!/bin/bash
CC=x86_64-linux-musl-gcc CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./bin/BingWallpaperServer -ldflags "-s -w" main.go
