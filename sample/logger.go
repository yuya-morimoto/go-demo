package sample

import (
	"io"
	"log"
	"os"
)

func Logger() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	flags := log.Lshortfile
	warnLogger := log.New(io.MultiWriter(file, os.Stderr), "WARN: ", flags)
	errorLogger := log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", flags)
	warnLogger.Println("warning A")

	errorLogger.Fatalln("critical error")
}
