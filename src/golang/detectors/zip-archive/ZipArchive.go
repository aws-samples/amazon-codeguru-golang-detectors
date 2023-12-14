package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// {fact rule=path-traversal@v1.0 defects=1}
func zipArchiveNoncompliant() error {
	archive := "PATH_TO_ZIP_FILE"
	target := "PATH_TO_EXTRACT"
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	for _, file := range reader.File {
		// Noncompliant: Extracting files from untrusted zip archives.
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, 0600)
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}
	return nil
}

// {/fact}

// {fact rule=path-traversal@v1.0 defects=0}
func zipArchiveCompliant() error {
	archive := "PATH_TO_ZIP_FILE"
	target := "PATH_TO_EXTRACT"
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	for _, file := range reader.File {
		// Compliant: Not Extracting files from untrusted zip archives.
		path := filepath.Join(target, "FILE")
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, 0600)
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		res := io.LimitReader(fileReader, 3)

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, res); err != nil {
			return err
		}
	}
	return nil
}

// {/fact}
