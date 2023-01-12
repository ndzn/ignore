package main

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Placeholder URL containing zip
	url := "https://www.bamsoftware.com/hacks/zipbomb/zbsm.zip"

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create("uwu.zip")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

	// Use os to unzip the file
	zipReader, err := zip.OpenReader("uwu.zip")
	if err != nil {
		panic(err)
	}
	defer zipReader.Close()

	// Iterate through files in zip
	for _, file := range zipReader.File {
		// Open file from zip
		zippedFile, err := file.Open()
		if err != nil {
			panic(err)
		}
		defer zippedFile.Close()

		// Get file info
		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		// Check the file is a directory
		if file.FileInfo().IsDir() {
			// Create directory
			os.MkdirAll(extractedFilePath, file.Mode())
			continue
		}

		// Create file
		outFile, err := os.OpenFile(
			extractedFilePath,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			file.Mode(),
		)
		if err != nil {
			panic(err)
		}
		defer outFile.Close()

		// Write file
		_, err = io.Copy(outFile, zippedFile)
		if err != nil {
			panic(err)
		}
	}
}