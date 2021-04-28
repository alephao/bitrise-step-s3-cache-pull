package main

import (
	"fmt"
	"os"

	"github.com/mholt/archiver"
)

func GenerateRestoreKeys(restoreKeys string) []string {
	checksumEngine := NewFileChecksumEngine()
	keyParser := NewKeyParser(&checksumEngine)

	var keys []string
	for _, keyTemplate := range keyParser.parseRestoreKeysInput(restoreKeys) {
		key := keyParser.parse(keyTemplate)
		keys = append(keys, key)
	}

	return keys
}

func main() {
	awsAccessKeyId := GetEnvOrExit("aws_access_key_id")
	awsSecretAccessKey := GetEnvOrExit("aws_secret_access_key")
	awsRegion := GetEnvOrExit("aws_region")
	bucketName := GetEnvOrExit("bucket_name")
	restoreKeys := GetEnvOrExit("restore_keys")
	cachePath := GetEnvOrExit("path")

	CreateTempFolder(func(tempFolderPath string) {
		s3 := NewAwsS3(
			awsRegion,
			awsAccessKeyId,
			awsSecretAccessKey,
			bucketName,
		)

		for _, key := range GenerateRestoreKeys(restoreKeys) {
			fmt.Printf("Checking if cache exists for key '%s'\n", key)
			cacheExists, cacheKey := s3.CacheExists(key)
			if cacheExists {
				fmt.Println("Cache found! Downloading...")
				downloadedFilePath := fmt.Sprintf("%s/%s.tar.gz", tempFolderPath, cacheKey)
				size, err := s3.Download(cacheKey, downloadedFilePath)

				if err != nil {
					fmt.Printf("Download failed with error: %s. Cancelling cache restore.\n", err.Error())
					return
				}

				fmt.Printf("Download was successful, file size: %d. Uncompressing...\n", size)

				err = archiver.Unarchive(downloadedFilePath, cachePath)

				if err != nil {
					fmt.Printf("Failed to uncompress: %s. Cancelling cache restore.\n", err.Error())
					return
				}
				return
			}
		}
		fmt.Println("Cache not found.")
	})

	os.Exit(0)
}
