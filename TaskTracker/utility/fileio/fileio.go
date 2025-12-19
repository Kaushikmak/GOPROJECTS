package fileio

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// get file path of data.json
func getDataFilePath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	exeDir := filepath.Dir(exePath)
	return filepath.Join(exeDir, "data", "data.json"), nil
}

// ReadData from file if not exits then create one
func ReadData() ([]byte, error) {
	path, err := getDataFilePath()
	if err != nil {
		return nil, err
	}
	// Create file if it doesn't exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Initialize with empty JSON array
		err := os.WriteFile(path, []byte("[]"), 0644)
		if err != nil {
			return nil, err
		}
		return []byte("[]"), nil
	}
	return os.ReadFile(path)
}

// WriteData overwrites the file with new data
func WriteData(data []byte) error {
	path, err := getDataFilePath()
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// AppendTask adds a serialized JSON object string to the JSON array in the file
func AppendTask(taskJSON string) error {
	// 1. Read existing data
	data, err := ReadData()
	if err != nil {
		return err
	}

	content := string(data)
	content = strings.TrimSpace(content)

	// 2. Handle empty file or empty list case
	if len(content) == 0 || content == "[]" {
		newContent := fmt.Sprintf("[%s]", taskJSON)
		return WriteData([]byte(newContent))
	}

	// 3. Validation: Ensure we are appending to a valid list
	if !strings.HasSuffix(content, "]") {
		return fmt.Errorf("file content is not a valid JSON array")
	}

	// 4. Modify String:
	// Remove the last ']'
	content = content[:len(content)-1]
	// Append ", " + newObject + "]"
	newContent := fmt.Sprintf("%s, %s]", content, taskJSON)

	// 5. Write back to file
	return WriteData([]byte(newContent))
}
