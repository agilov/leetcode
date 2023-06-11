package helpers

import (
	"encoding/json"
	"io"
	"os"
)

// Function that takes filename with json data and return array of testcases R
func TestDataProvider[R any](filename string) ([]R, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var result []R

	err = json.Unmarshal(data, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
