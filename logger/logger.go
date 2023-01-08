package logger

import (
	"log"
	"os"
)

func InitLogger(fpath string) {
	fpLog, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(fpLog)
}
