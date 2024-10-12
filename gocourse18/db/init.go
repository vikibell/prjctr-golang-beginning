package db

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/vikibell/prjctr-golang-beginning/gocourse18/domains/user/model"
	"io"
	"os"
	"path/filepath"
)

var schema = `
CREATE TABLE if not exists users
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
);
`

func getUsersFromFile() []model.User {
	var users []model.User

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

func Init() *sqlx.DB {
	dsn := "vika:123@tcp(127.0.0.1:3320)/db?parseTime=true"

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("Failed to connect database")
	}

	db.MustExec(schema)

	return db
}

func PopulateUsers(db *sqlx.DB) {
	query := `TRUNCATE TABLE users` // For other DBs (Postgres, MySQL)
	_, trancErr := db.Exec(query)
	if trancErr != nil {
		fmt.Printf("Failed to truncate table: %v\n", trancErr)
		return
	}

	users := getUsersFromFile()

	tx := db.MustBegin()
	for _, user := range users {
		_, insertErr := tx.NamedExec(`
		INSERT INTO users (name, surname, email, age, sex, city, taxi_count, profession) 
		VALUES (:name, :surname, :email, :age, :sex, :city, :taxi_count, :profession)`,
			&user)
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
