#!/usr/bin/env bash

path_val=$(pwd)
app_name=${path_val##*/}

rm -rf bin
mkdir bin

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
mv ${app_name} ${app_name}-linux-amd64
mv ${app_name}-linux-amd64 ./bin

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
mv ${app_name} ${app_name}-darwin-amd64
mv ${app_name}-darwin-amd64 ./bin

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
mv ${app_name}.exe ${app_name}-windows-amd64.exe
mv ${app_name}-windows-amd64.exe ./bin