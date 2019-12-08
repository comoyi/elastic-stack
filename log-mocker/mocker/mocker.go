package mocker

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

func Start(filePath string, count int) {
	var logData string
	var file *os.File
	var err error
	var dir string
	dir = path.Dir(filePath)
	if _, err = os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(dir, 0755); err != nil {
				fmt.Println("Create dir failed", err)
				return
			}
		}
	}

	file, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Open file failed", err)
		return
	}
	defer file.Close()

	var line int
	var i = 0
	var dateTime string
	var appFiles []string
	appFiles = []string{
		"/path/to/app/user.go",
		"/path/to/app/banner.go",
		"/path/to/app/log.go",
		"/path/to/app/db/postgresql.go",
		"/path/to/app/db/mysql.go",
	}
	var appFile string
	var errorLevels []string
	errorLevels = []string{
		"DEBUG",
		"INFO",
		"WARNING",
		"ERROR",
		"FATAL",
	}
	var errorLevel string
	var logContent string
	rand.Seed(time.Now().UnixNano())
	for {
		if count != 0 {
			i++
		}
		if count != 0 && i > count {
			//fmt.Println("Reach loop limit")
			break
		}
		dateTime = time.Now().Format("2006-01-02 15:04:05")
		appFile = appFiles[rand.Intn(5)]
		line = rand.Intn(1777)
		errorLevel = errorLevels[rand.Intn(5)]
		logContent = "log content random: " + strconv.Itoa(rand.Intn(99999))
		logData = fmt.Sprintf("%s %s %d %s %s\n", dateTime, appFile, line, errorLevel, logContent)
		if _, err := file.WriteString(logData); err != nil {
			fmt.Println("Write log failed", err)
			break
		}
		<-time.After(time.Millisecond * time.Duration(rand.Intn(300)+100))
	}

	//fmt.Println("Mock finished")
}
