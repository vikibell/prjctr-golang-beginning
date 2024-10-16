package service

import (
	"encoding/json"
	"github.com/vikibell/prjctr-golang-beginning/gocourse18/domains/user"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

func GetUsersFromFile() []user.User {
	var users []user.User

	currentDir, err := os.Getwd()
	if err != nil {
		slog.Error("Failed to get current directory", "error", err)
		return users
	}

	file, errOpen := os.Open(filepath.Join(currentDir, "/assets/users_another.json"))
	if errOpen != nil {
		slog.Error("Failed to open file", "error", err)
		return users
	}

	defer file.Close()

	byteValue, errRead := io.ReadAll(file)
	if errRead != nil {
		slog.Error("Failed to read file", "error", errRead)
		return users
	}

	errUnm := json.Unmarshal(byteValue, &users)
	if errUnm != nil {
		slog.Error("Failed to unmarshal json", "error", errRead)
		return users
	}

	return users
}
