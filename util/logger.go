package util

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	// Err 记录错误日志
	Err *log.Logger
	// Warn 记录警告日志
	Warn *log.Logger
	// Inform 记录信息日志
	Inform *log.Logger
	// Debug 记录Debug日志
	Debug *log.Logger
)

var file *os.File

func init() {
	f, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file: ", err)
	}
	file = f
}

// SetLogLevel 设置日志级别
func SetLogLevel(level string) {
	switch level {
	case "error":
		Err = log.New(io.MultiWriter(file, os.Stderr), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warn = log.New(ioutil.Discard, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
		Inform = log.New(ioutil.Discard, "Information: ", log.Ldate|log.Ltime|log.Lshortfile)
		Debug = log.New(ioutil.Discard, "Debugging: ", log.Ldate|log.Ltime|log.Lshortfile)
	case "warning":
		Err = log.New(io.MultiWriter(file, os.Stderr), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warn = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
		Inform = log.New(ioutil.Discard, "Information: ", log.Ldate|log.Ltime|log.Lshortfile)
		Debug = log.New(ioutil.Discard, "Debugging: ", log.Ldate|log.Ltime|log.Lshortfile)
	case "information":
		Err = log.New(io.MultiWriter(file, os.Stderr), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warn = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
		Inform = log.New(os.Stdout, "Information: ", log.Ldate|log.Ltime|log.Lshortfile)
		Debug = log.New(ioutil.Discard, "Debugging: ", log.Ldate|log.Ltime|log.Lshortfile)
	case "debugging":
		Err = log.New(io.MultiWriter(file, os.Stderr), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warn = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
		Inform = log.New(os.Stdout, "Information: ", log.Ldate|log.Ltime|log.Lshortfile)
		Debug = log.New(os.Stdout, "Debugging: ", log.Ldate|log.Ltime|log.Lshortfile)
	default:
		Err = log.New(io.MultiWriter(file, os.Stderr), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warn = log.New(ioutil.Discard, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
		Inform = log.New(ioutil.Discard, "Information: ", log.Ldate|log.Ltime|log.Lshortfile)
		Debug = log.New(ioutil.Discard, "Debugging: ", log.Ldate|log.Ltime|log.Lshortfile)
	}
}
