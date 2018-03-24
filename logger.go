package logger

import (
	"os"
	"log"
	"time"
)

var l       = log.New(os.Stdout, "", 0)
var dirSave = ""

func SaveDir(dir string) {
	dirSave = dir
}

func Info (str string) {
	if dirSave != "" {
		writeLog("[INFO]: " + str)
	}
	l.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " [INFO]: ")
	l.Print(str)
}

func Error (str string) {
	if dirSave != "" {
		writeLog("[ERROR]: " + str)
	}
	l.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " [ERROR]: ")
	l.Print(str)
}

func writeLog(str string) {
	filename := dirSave + "/" + time.Now().Format("20060102") + ".log"
	f, err := os.OpenFile(filename, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		dirSave = ""
    	log.Printf("[ERROR]: can't write log file : %v", err)
	}else{
		log.SetOutput(f)
		log.Println(str);
	}
}