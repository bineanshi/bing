#!/bin/bash
CC=x86_64-linux-musl-gcc CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"  -o ./bin/BingWallpaperServer