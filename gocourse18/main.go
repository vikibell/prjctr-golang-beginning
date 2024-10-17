package main

import (
	"github.com/vikibell/prjctr-golang-beginning/gocourse18/domains/statistic"
	"github.com/vikibell/prjctr-golang-beginning/gocourse18/domains/user"
	"github.com/vikibell/prjctr-golang-beginning/gocourse18/mysql"
	"log/slog"
)

func main() {
	database := mysql.NewDB(mysql.DSN).Connection

	usersRepository := user.NewUserRepository(database)
	users, totalUsers, err := usersRepository.FindAll(1, 10)
	if err != nil {
		slog.Error("Failed to find users", "error", err)
		return
	}

	for _, selectedUser := range users {
		slog.Info("User", "details", selectedUser)
	}
	slog.Info("Total users", "count", totalUsers)

	statisticRepository := statistic.NewStatisticsRepository(database)
	statisticsList, totalStat, err := statisticRepository.FindAll(1, 10)
	if err != nil {
		slog.Error("Failed to find statistics", "error", err)
		return
	}

	for _, stat := range statisticsList {
		slog.Info("Statistics", "details", stat)
	}
	slog.Info("Total statistic records", "count", totalStat)
}
