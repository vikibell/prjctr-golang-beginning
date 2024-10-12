package db

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/vikibell/prjctr-golang-beginning/gocourse18/domains/statistic"
	"github.com/vikibell/prjctr-golang-beginning/gocourse18/domains/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"os"
	"path/filepath"
)

const dsn = "vika:123@tcp(127.0.0.1:3320)/db?parseTime=true"

var usersSchema = `
CREATE TABLE IF NOT EXISTS users
(
    id         bigint unsigned auto_increment primary key,
    name       longtext         not null,
    surname    longtext         not null,
    email      longtext         not null,
    age        bigint default 0 not null,
    sex        longtext         not null,
    city       longtext         not null,
    taxi_count bigint default 0 null,
    profession longtext         null
);`

var statisticsSchema = `
CREATE TABLE IF NOT EXISTS statistics
(   
    id            bigint unsigned auto_increment primary key,
    city          longtext         not null,
    average_trips bigint default 0 null,
    age_range     int    default 0 null
);`

func init() {
	populateDB()
}

func GetConnection() *gorm.DB {
	gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return gormdb
}

func getUsersFromFile() []user.User {
	var users []user.User

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get current directory: %v\n", err)
		return users
	}

	file, errOpen := os.Open(filepath.Join(currentDir, "/db/users_another.json"))
	if errOpen != nil {
		fmt.Printf("Failed to open file: %v\n", errOpen)
		return users
	}

	defer file.Close()

	byteValue, errRead := io.ReadAll(file)
	if errRead != nil {
		fmt.Printf("Failed to read file: %v\n", errRead)
		return users
	}

	errUnm := json.Unmarshal(byteValue, &users)
	if errUnm != nil {
		fmt.Printf("Failed to unmarshal json: %v\n", errUnm)
		return users
	}

	return users
}

func populateDB() {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("Failed to connect database")
	}

	db.MustExec(usersSchema)
	populateUsers(db)
	db.MustExec(statisticsSchema)
	populateStatistics(db)
	db.Close()
}

func populateUsers(db *sqlx.DB) {
	query := `TRUNCATE TABLE users`
	_, trancErr := db.Exec(query)
	if trancErr != nil {
		fmt.Printf("Failed to truncate table: %v\n", trancErr)
		return
	}

	users := getUsersFromFile()

	tx := db.MustBegin()
	_, insertErr := tx.NamedExec(`
		INSERT INTO users (name, surname, email, age, sex, city, taxi_count, profession) 
		VALUES (:name, :surname, :email, :age, :sex, :city, :taxi_count, :profession)`,
		users)
	if insertErr != nil {
		fmt.Printf("Failed to insert into table: %v\n", insertErr)
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		fmt.Printf("Failed to commit: %v\n", commitErr)
		return
	}
}

func populateStatistics(db *sqlx.DB) {
	query := `TRUNCATE TABLE statistics`
	_, trancErr := db.Exec(query)
	if trancErr != nil {
		fmt.Printf("Failed to truncate table: %v\n", trancErr)
		return
	}

	var cities []string
	err := db.Select(&cities, "SELECT DISTINCT city FROM users")
	if err != nil {
		fmt.Printf("Failed to select: %v\n", err)
		return
	}

	tx := db.MustBegin()
	for _, city := range cities {
		_, insertErr := tx.NamedExec(`INSERT INTO statistics (city) VALUES (:city)`, statistic.Statistic{City: city})
		if insertErr != nil {
			fmt.Printf("Failed to insert into table: %v\n", insertErr)
		}
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		fmt.Printf("Failed to commit: %v\n", commitErr)
		return
	}
}
