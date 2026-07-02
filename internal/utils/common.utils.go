package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func GetCurrentDir() (string, error) {
	ex, err := os.Executable()

	if err != nil {
		return "", err
	}

	return filepath.Dir(ex), nil
}
