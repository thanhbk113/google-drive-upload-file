package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"thanhbk113/internal/config"
)

func InitFileGoogleDriverAuth() {
	fileURL := config.GetEnv().GOOGLE_DRIVER_AUTH

	dir, _ := os.Getwd()
	outputPath := path.Join(dir + config.PathGoogleDriverAuth)
	err := downloadFile(fileURL, outputPath)
	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
	} else {
		fmt.Println("File downloaded successfully.")
	}

	fmt.Println("Init file GoogleDriverAuth Successfully!")
}

func downloadFile(url, outputPath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	file.Close()

	return nil
}
