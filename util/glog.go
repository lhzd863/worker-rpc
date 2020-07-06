package util

import (
	"log"
	"os"
	"regexp"
	"time"
)

func Glog(f string, context string) {
	timeStr := time.Now().Format("20060102")
	dt := "\\$\\{"+ENV_VAR_DATE+"\\}"
	reg := regexp.MustCompile(dt)
	f1 := reg.ReplaceAllString(f, timeStr)
	logFile, err := os.OpenFile(f1, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	defer logFile.Close()
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.SetPrefix("[r]")
	log.SetFlags(log.LstdFlags | log.Ltime)
	log.Println(context)
	return
}
