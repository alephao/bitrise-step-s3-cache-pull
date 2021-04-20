package main

import (
	"fmt"
	"os"
)

func CreateTempFolder(f func(tempFolderPath string)) {
	path := "~/bitrise-s3-step-pull-tmp"
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f(path)

	err = os.RemoveAll(path)

	if err != nil {
		fmt.Println(err)
	}
}

func GetEnvOrExit(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Printf("Missing variable '%s'\n", key)
		os.Exit(1)
	}
	return value
}
