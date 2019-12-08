package main

import (
	"flag"
	"fmt"
	"github.com/comoyi/log-mocker/mocker"
	"time"
)

func main() {
	var count int
	var filePath string
	flag.IntVar(&count, "c", 100, "Number of simulated logs, if it is 0, there is no limit")
	flag.StringVar(&filePath, "f", "", "Log file path (default: /tmp/log/app/<date +%Y-%m-%d>/app.log)")
	flag.Parse()
	if filePath == "" {
		var dateTime string
		dateTime = time.Now().Format("2006-01-02")
		filePath = fmt.Sprintf("/tmp/log/app/%s/app.log", dateTime)
	}
	mocker.Start(filePath, count)
}
