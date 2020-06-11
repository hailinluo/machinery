package log

import (
	"os"
	"path/filepath"
)

func GetLogFilePath(logName string) string {
	file, err := filepath.Abs(os.Args[0])
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(file)
	pDir := filepath.Dir(dir + "..")
	logDir := pDir + "/log/"
	if !exist(logDir) {
		if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
			panic(err)
		}
	}

	return logDir + logName + ".log"
}

func exist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			return true
		}
	}
	return false
}
