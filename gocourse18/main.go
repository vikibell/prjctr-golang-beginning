package main

import (
	"fmt"
	"github.com/vikibell/prjctr-golang-beginning/gocourse18/db"
	"github.com/vikibell/prjctr-golang-beginning/gocourse18/domains/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	database := db.Init()
	db.PopulateUsers(database)
	database.Close()

	dsn := "vika:123@tcp(127.0.0.1:3320)/db?parseTime=true"
	gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	usersRepository := user.NewUserRepository(gormdb)
	users, total, err := usersRepository.FindAll(1, 10)
	if err != nil {
		return
	}
	for _, selectedUser := range users {
		fmt.Printf("User: %v\n", selectedUser)
	}
	fmt.Printf("Total users: %v\n", total)
}
