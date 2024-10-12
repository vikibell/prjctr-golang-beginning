package main

import (
	"fmt"
	"github.com/vikibell/prjctr-golang-beginning/gocourse18/db"
	"github.com/vikibell/prjctr-golang-beginning/gocourse18/domains/statistic"
	"github.com/vikibell/prjctr-golang-beginning/gocourse18/domains/user"
)

func main() {
	database := db.GetConnection()

	usersRepository := user.NewUserRepository(database)
	users, totalUsers, err := usersRepository.FindAll(1, 10)
	if err != nil {
		return
	}

	for _, selectedUser := range users {
		fmt.Printf("User: %v\n", selectedUser)
	}
	fmt.Printf("Total users: %v\n", totalUsers)

	statisticRepository := statistic.NewStatisticsRepository(database)
	statisticsList, totalStat, errStatFetch := statisticRepository.FindAll(1, 10)
	if errStatFetch != nil {
		return
	}

	for _, stat := range statisticsList {
		fmt.Printf("Statistics: %v\n", stat)
	}
	fmt.Printf("Total statistic records: %v\n", totalStat)
}
