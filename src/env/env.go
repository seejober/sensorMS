package env

import (
	"os"
	"runtime"
)

var ostype = runtime.GOOS

func getProjectPath() string {
	var projectPath string
	projectPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return projectPath
}

func GetConfigPath() string {
	path := getProjectPath()
	if ostype == "windows" {
		path = path
	} else if ostype == "linux" {
		path = path + "/" + ".."
	}
	return path
}
