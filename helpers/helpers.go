package helpers

import (
	"os"
)

func ReadFile(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}
