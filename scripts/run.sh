#!/bin/sh
srcPath="cmd"
pkgFile="main.go"
outputPath="build"
app="poker-server"
output="$outputPath/$app"
src="$srcPath/$app/$pkgFile"

printf "\nRunning app: $app\n"
time go run $src
printf "\nStopping app: $app\n\n"