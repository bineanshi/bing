#!/bin/bash
CC=x86_64-linux-musl-gcc CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o dist/snatch -ldflags "-linkmode external -extldflags -static" cron.go
CC=x86_64-linux-musl-gcc CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o dist/BingWallpaperServer -ldflags "-linkmode external -extldflags -static" main.go
