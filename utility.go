package main

import (
	"log"
	"os"
)

func CreateTempFolder(f func(tempFolderPath string)) {
	path := "/tmp/bitrise-s3-step-pull-tmp"
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		log.Fatalln(err.Error())
	}

	f(path)

	err = os.RemoveAll(path)

	if err != nil {
		log.Println(err.Error())
	}
}

func GetEnvOrExit(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing variable '%s'\n", key)
	}
	return value
}
