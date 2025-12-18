package fileio

import (
	"os"
	"path/filepath"
)

// function to get file path of data.json
func getDataFilePath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	exeDir := filepath.Dir(exePath)

	dataPath := filepath.Join(exeDir, "data", "data.json")

	return dataPath, nil
}

// function to read an existing file
func ReadData() ([]byte, error) {
	path, err := getDataFilePath()
	if err != nil {
		return nil, err
	}
	return os.ReadFile(path)
}

// function to append data to file
// func AppendFile() ([]byte, error)

// function to delete file
// func DeleteFile()
