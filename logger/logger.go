package logger

import (
	"log"
	"os"
)

var (
	warning   *log.Logger
	info      *log.Logger
	error     *log.Logger
	debugging *log.Logger
)

func Init() {
	info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Warning(message string){
	info.Println(message)
}

func Info(message string){
	info.Println(message)
}

func Error(message string){
	error.Println(message)
}
