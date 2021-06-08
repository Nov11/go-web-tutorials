package pkg

import (
	"log"
	"os"
)

func FileLogger() *log.Logger {
	f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	return log.New(f, "INFO\t", log.Ldate|log.Ltime)
}
