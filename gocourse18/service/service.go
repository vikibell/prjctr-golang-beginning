package service

import (
	"encoding/json"
	"github.com/vikibell/prjctr-golang-beginning/gocourse18/domains/user"
	"io"
	"os"
	"path/filepath"
)

func GetUsersFromFile(path, filename string) ([]user.User, error) {
	var users []user.User

	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(currentDir, path, filename))
	if err != nil {
		return nil, err
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
